package tonseapi

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
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
}


func graphql() {
    os.Chdir(tonossePath+"/graphql/package")
    godotenv.Load()
    cmd := exec.Command(tonossePath+"/graphql/nodejs/bin/node", "index.js")
    if runtime.GOOS == "darwin" {
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
}

func nginx() {
    cmd := exec.Command("nginx -g 'daemon on; master_process on;'")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Start()
}

var PIDFile = "./.daemonize.pid"
func stop() {
    if _, err := os.Stat(PIDFile); err == nil {
        data, err := ioutil.ReadFile(PIDFile)
        if err != nil {
            log.Fatal("Not running")
            os.Exit(1)
        }
        ProcessID, err := strconv.Atoi(string(data))

        if err != nil {
            log.Fatal("Unable to read and parse process id found in ", PIDFile)
            os.Exit(1)
        }
        process, err := os.FindProcess(ProcessID)
        log.Printf(string(ProcessID))
        if err != nil {
            log.Fatal("Unable to find process ID [%v] with error %v \n", ProcessID, err)
            os.Exit(1)
        }
        // remove PID file
        os.Remove(PIDFile)

        log.Printf("Killing process ID [%v] now.\n", ProcessID)
        // kill process and exit immediately
        err = process.Kill()
        if err != nil {
            log.Fatal("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
            os.Exit(1)
        } else {
            log.Printf("Killed process ID [%v]\n", ProcessID)
            os.Exit(0)
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
