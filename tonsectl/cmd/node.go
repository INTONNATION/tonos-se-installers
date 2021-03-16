package cmd

import (
    "fmt"
    "os/exec"
    //"path/filepath"
    "log"
    //"runtime"
    "github.com/spf13/cobra"
    "compress/gzip"
    "net/http"
    "io"
    "os"
    "archive/tar"
)

func init() {
  rootCmd.AddCommand(nodeCmd)
}

var tonosseUrl = "https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-0.25.0/"
var tonosseTar = "tonos-se-v-0.25.0.tgz"
var tonossePath = "/opt/tonsectl/"
var tonosseConfigUrl = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/cfg"
var tonosseLogCfg = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/log_cfg.yml"
var tonossePrivKey = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/private-key"
var tonossePubKey = "https://raw.githubusercontent.com/tonlabs/tonos-se/master/docker/ton-node/pub-key"

var nodeCmd = &cobra.Command{
    Use:   "node",
    Short: "Start TON node service",
    Run: func(cmd *cobra.Command, args []string) {
       node()
    },

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

