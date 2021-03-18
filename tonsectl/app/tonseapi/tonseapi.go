package tonseapi

import (
    "fmt"
    "os/exec"
    "log"
    "compress/gzip"
    "net/http"
    "syscall"
    "time"
    //"github.com/gorilla/mux"
    "io"
    "os"
    "os/user"
    "archive/tar"
)

var tonosseUrl = "https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-0.25.0/"
var tonosseTar = "tonos-se-v-0.25.0.tgz"
var tonossePath = "/tonse"
var tonosseConfigUrl = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/cfg"
var tonosseLogCfg = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/log_cfg.yml"
var tonossePrivKey = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/private-key"
var tonossePubKey = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/pub-key"

var pid = 0

func tonseapi() {
    //myRouter := mux.NewRouter().StrictSlash(true)
    http.HandleFunc("/tonse/init", tonseInit)
    http.HandleFunc("/tonse/start", tonseStart)
    http.HandleFunc("/tonse/stop", tonseStop)
    http.HandleFunc("/tonse/status", tonseStatus)
    http.HandleFunc("/tonse/reset", tonseReset)
    http.HandleFunc("/tonse/upgrade", tonseUpgrade)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func tonseInit(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseInit")
}

func tonseStart(w http.ResponseWriter, r *http.Request){
    //node()
    arangodStart()
    fmt.Println("Endpoint Hit: tonseStart")
}

func tonseStop(w http.ResponseWriter, r *http.Request){
    arangodStop()
    fmt.Println("Endpoint Hit: tonseStop")
}

func tonseStatus(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseStatus")
}

func tonseReset(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseReset")
}

func tonseUpgrade(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: tonseUpgrade")
}

func downloadFile(filepath string, url string) (err error) {

  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()

  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("bad status: %s", resp.Status)
  }

  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

func extractTarGz(gzipStream io.Reader) {
    uncompressedStream, err := gzip.NewReader(gzipStream)
    if err != nil {
        log.Fatal("extractTarGz: NewReader failed")
    }

    tarReader := tar.NewReader(uncompressedStream)

    for true {
        header, err := tarReader.Next()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatalf("extractTarGz: Next() failed: %s", err.Error())
        }

        switch header.Typeflag {
        case tar.TypeDir:
            if err := os.Mkdir(header.Name, 0755); err != nil {
                log.Fatalf("extractTarGz: Mkdir() failed: %s", err.Error())
            }
        case tar.TypeReg:
            outFile, err := os.Create(tonossePath + header.Name)
            if err != nil {
                log.Fatalf("extractTarGz: Create() failed: %s", err.Error())
            }
            if _, err := io.Copy(outFile, tarReader); err != nil {
                log.Fatalf("extractTarGz: Copy() failed: %s", err.Error())
            }
            outFile.Close()

        default:
            log.Fatalf(
                "extractTarGz: uknown type: %s in %s",
                header.Typeflag,
                header.Name)
        }
    }
}

func node() {
    //Commented because of private repo
    //downloadFile(tonossePath+tonosseTar, tonosseUrl+tonosseTar)
    usr, e := user.Current()
    if e != nil {
        log.Fatal( e )
    }
    fmt.Println( usr.HomeDir )
    tarFile, err1 := os.Open(usr.HomeDir + tonossePath + tonosseTar)
    if err1 != nil {
        log.Fatalf(err1.Error())
    }
    extractTarGz(tarFile)
    downloadFile(usr.HomeDir + tonossePath+"cfg", tonosseConfigUrl)
    downloadFile(usr.HomeDir + tonossePath+"log_cfg.yml", tonosseLogCfg)
    downloadFile(usr.HomeDir + tonossePath+"private-key", tonossePrivKey)
    downloadFile(usr.HomeDir + tonossePath+"pub-key", tonossePubKey)
    os.Chdir(usr.HomeDir + tonossePath)
    os.Chmod("ton_node_startup", 0700)
    cmd := exec.Command("./ton_node_startup", "--config", "cfg")
    cmd.Start()
}

func arangodStop(){
   err := syscall.Kill(pid, 9)
   if err == nil {
      fmt.Println("Signal SIGTERM sent to PID", pid)
   }
}

func arangodStart(){
        os.Chdir(tonossePath+"/arangodb/usr/bin")
        usr, e := user.Current()
        if e != nil {
            log.Fatal( e )
        }
        fmt.Println( usr.HomeDir )
	upgrade := exec.Command("arangod", "--config", usr.HomeDir + tonossePath + "/arangodb/etc/config", "--server.endpoint", "tcp://127.0.0.1:8529", "--server.authentication=false", "--log.foreground-tty", "true", "--database.auto-upgrade", "true")
        upgrade.Stdout = os.Stdout
	upgrade.Stderr = os.Stderr
	err := upgrade.Run()
	if err != nil {
            log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	cmd := exec.Command("arangod", "--config", usr.HomeDir + tonossePath + "/arangodb/etc/config", "--server.endpoint", "tcp://127.0.0.1:8529", "--server.authentication=false", "--log.foreground-tty", "true")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
	cmd.Start()
	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	pid = cmd.Process.Pid
	for {
	    status := exec.Command("arangosh", "--server.endpoint=127.0.0.1", "--server.authentication=false", "--javascript.execute-string", "'db._version()'")
	    status.Stdout = os.Stdout
	    status.Stderr = os.Stderr
	    err := status.Run()
	    time.Sleep(1 * time.Second)
	    if err == nil {
	        break
	    }
	}
	dump := exec.Command("arangosh", "--server.authentication", "false", "--server.endpoint=tcp://127.0.0.1:8529", "--javascript.execute", usr.HomeDir + tonossePath + "/arangodb/initdb.d/upgrade-arango-db.js")
	dump.Stdout = os.Stdout
	dump.Stderr = os.Stderr
	dump.Run()
}

func RunApi() {
    tonseapi()
}
