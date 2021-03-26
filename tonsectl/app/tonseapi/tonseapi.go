package tonseapi

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/mitchellh/go-ps"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "os/user"
    "path"
    "runtime"
    "strconv"
    "time"
)


var usr, e = user.Current()
var tonossePath = usr.HomeDir + "/tonse/"

var pid = 0
var PIDFile = tonossePath+".daemonize.pid"


func tonseapi() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/tonse/start", tonseStart)
    myRouter.HandleFunc("/tonse/stop", tonseStop)
    myRouter.HandleFunc("/tonse/status", tonseStatus)
    myRouter.HandleFunc("/tonse/reset", tonseReset)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}


func tonseStart(w http.ResponseWriter, r *http.Request){
    arangodStart()
    node()
    graphql()
    nginx()
    fmt.Println("Endpoint Hit: tonseStart")
}

func tonseStop(w http.ResponseWriter, r *http.Request){
    log.Println("Endpoint Hit:1")
    stop()
    log.Println("Endpoint Hit: tonseStop")
}

func tonseStatus(w http.ResponseWriter, r *http.Request){
    status()
    fmt.Println("Endpoint Hit: tonseStatus")
}

func tonseReset(w http.ResponseWriter, r *http.Request){
    reset_dir()
    fmt.Println("Endpoint Hit: tonseReset")
}

func node() {
    os.Chdir(tonossePath + "/node")
    cmd := exec.Command("./ton_node_startup", "--config", "cfg")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Start()
    cmd.Wait()
}


func arangodStart(){
        os.Chdir(tonossePath+"/arangodb/bin")
	upgrade := exec.Command("./arangod", "--config", tonossePath + "/arangodb/etc/config", "--server.endpoint", "tcp://127.0.0.1:8529", "--server.authentication=false", "--log.foreground-tty", "true", "--database.auto-upgrade", "true")
        upgrade.Stdout = os.Stdout
	upgrade.Stderr = os.Stderr
	err := upgrade.Run()
	if err != nil {
            log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	cmd := exec.Command("./arangod", "--config", tonossePath + "/arangodb/etc/config", "--server.endpoint", "tcp://127.0.0.1:8529", "--server.authentication=false", "--log.foreground-tty", "true")
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
	cmd.Start()
	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	pid = cmd.Process.Pid
	for {
	    status := exec.Command("./arangosh", "--server.endpoint=127.0.0.1", "--server.authentication=false", "--javascript.execute-string", "'db._version()'")
	    status.Stdout = os.Stdout
	    status.Stderr = os.Stderr
	    err := status.Run()
	    time.Sleep(1 * time.Second)
	    if err == nil {
	        break
	    }
	}
	dump := exec.Command("./arangosh", "--server.authentication", "false", "--server.endpoint=tcp://127.0.0.1:8529", "--javascript.execute", tonossePath + "/arangodb/initdb.d/upgrade-arango-db.js")
	dump.Stdout = os.Stdout
	dump.Stderr = os.Stderr
	dump.Run()
    dump.Wait()
}


func graphql() {
    os.Chdir(tonossePath+"/graphql/package")
    godotenv.Load()
    cmd := exec.Command(tonossePath+"/graphql/nodejs/bin/node", "index.js")
    if runtime.GOOS == "darwin" {
        cmd = exec.Command("node", "index.js")
    }
    if runtime.GOOS == "linux" {
        cmd = exec.Command("node", "index.js")
    }
    if runtime.GOOS == "windows" {
        cmd = exec.Command("node", "index.js")
    }
    f, err := os.OpenFile("./APIlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }
    defer f.Close()
    // On this line you're going to redirect the output to a file
    cmd.Stdout = f
    cmd.Start()
    cmd.Wait()
}

func nginx() {
    cmd := exec.Command("nginx -g 'daemon on; master_process on;'")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Start()
}



func stop() {
    data, err := ioutil.ReadFile(PIDFile)
    if err != nil {
        log.Fatal("Not running")
        os.Exit(1)
    }
    list, err := ps.Processes()
    if err != nil {
        panic(err)
    }
    ProcessID, err := strconv.Atoi(string(data))
    for _, p := range list {
        if p.PPid() == ProcessID {
            process, _ := os.FindProcess(p.Pid())
            log.Printf("Process %s with PID %d and PPID %d", p.Executable(), p.Pid(), p.PPid())
            process.Kill()
        }
    }
}

func status() {
    data, err := ioutil.ReadFile(PIDFile)
    if err != nil {
        log.Fatal("Not running")
        os.Exit(1)
    }
    list, err := ps.Processes()
    if err != nil {
        panic(err)
    }
    ProcessID, err := strconv.Atoi(string(data))
    for _, p := range list {
        if p.PPid() == ProcessID {
            log.Printf("Process %s with PID %d and PPID %d", p.Executable(), p.Pid(), p.PPid())
        }
    }
}


func reset_dir()  {
    dir, err := ioutil.ReadDir(tonossePath)
    if err != nil {
     fmt.Print("Cant find TONSE dir")
    }
    for _, d := range dir {
        os.RemoveAll(path.Join([]string{"tmp", d.Name()}...))
    }
}
func main(){
    lf, err := os.OpenFile("./APIlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer lf.Close()
    log.SetOutput(lf)
    tonseapi()
}
func RunApi() {
    main()
}
