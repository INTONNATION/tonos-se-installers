package tonseapi

import (
    "fmt"
    "os/exec"
    "log"
    "compress/gzip"
    "net/http"
    //"github.com/gorilla/mux"
    "io"
    "os"
    "archive/tar"
)

var tonosseUrl = "https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-0.25.0/"
var tonosseTar = "tonos-se-v-0.25.0.tgz"
var tonossePath = "/opt/tonsectl/"
var tonosseConfigUrl = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/cfg"
var tonosseLogCfg = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/log_cfg.yml"
var tonossePrivKey = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/private-key"
var tonossePubKey = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/pub-key"

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
    node()
    fmt.Println("Endpoint Hit: tonseStart")
}

func tonseStop(w http.ResponseWriter, r *http.Request){
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
    tarFile, err1 := os.Open(tonossePath + tonosseTar)
    if err1 != nil {
        log.Fatalf(err1.Error())
    }
    extractTarGz(tarFile)
    downloadFile(tonossePath+"cfg", tonosseConfigUrl)
    downloadFile(tonossePath+"log_cfg.yml", tonosseLogCfg)
    downloadFile(tonossePath+"private-key", tonossePrivKey)
    downloadFile(tonossePath+"pub-key", tonossePubKey)
    os.Chdir(tonossePath)
    os.Chmod("ton_node_startup", 0700)
    cmd := exec.Command("./ton_node_startup", "--config", "cfg")
    cmd.Start()
}

func RunApi() {
    tonseapi()
}
