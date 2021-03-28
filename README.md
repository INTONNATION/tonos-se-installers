## Introduction

[TON OS Startup Edition installers contest for different platforms](https://forum.freeton.org/t/ton-os-startup-edition-installers-contest-for-different-platforms/9101) has several main goals:



*   non-docker installation 
*   API and CLI for management
*   automatic builds when new release published
*   quick and easy to use solution
*   compatible with all popular Operating Systems

That’s why our team developed a cross platform solution which was verified on Windows 10, Mac  OS Mojave/Catalina, Ubuntu 18.04/20.04.


## System requirements


<table>
  <tr>
   <td><strong>Configuration</strong>
   </td>
   <td><strong>CPU (threads)</strong>
   </td>
   <td><strong>RAM (GiB)</strong>
   </td>
   <td><strong>Storage (GiB)</strong>
   </td>
   <td><strong>Operating system</strong>
   </td>
  </tr>
  <tr>
   <td>Recommended
   </td>
   <td>4
   </td>
   <td>8
   </td>
   <td>50
   </td>
   <td>Windows 10, Mac  OS Mojave/Catalina, Ubuntu 18.04/20.04
   </td>
  </tr>
</table>

## Quick start

### Windows

1. Open CMD
2. Download the latest of TON OS SE from our [Github Releases](https://github.com/INTONNATION/tonos-se-installers/releases):

       curl -o tonsectl.exe https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-0.25.0/tonsectl_windows.exe

3. Install required dependencies:

       tonsectl.exe install

4. Start TON OS SE:

       tonsectl.exe start

### Linux

1. Download the latest of TON OS SE from our [Github Releases](https://github.com/INTONNATION/tonos-se-installers/releases):

        curl -LJ -o tonsectl https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-0.25.0/tonsectl_linux

2. Make tonsectl executable:

       chmod +x tonsectl

3. Install required dependencies:

       ./tonsectl install

4. Start TON OS SE:

       ./tonsectl start

### OSX


1. Download the latest of TON OS SE from our [Github Releases](https://github.com/INTONNATION/tonos-se-installers/releases):

        curl -LJ -o tonsectl https://github.com/INTONNATION/tonos-se-installers/releases/download/tonos-se-v-0.25.0/tonsectl_darwin

2. Make tonsectl executable:

       chmod +x tonsectl

3. Install required dependencies:

        ./tonsectl install

4. Start TON OS SE:

       ./tonsectl start



## Project description

Based on contest goals we choose the following architecture and tools:



*   cross platform open source programming language (GO lang) which allows to build an app as a simple binary which works on all Operating Systems
*   In our case this binary is called [tonsectl](https://github.com/INTONNATION/tonos-se-installers/tree/master/tonsectl). It allows users to manage TON OS SE with CLI approach as well as spin up an API for TON OS SE management.
*   In order to support Continuous Integration with TON Labs [tonos-se repository](https://github.com/tonlabs/tonos-se) and automatically get the latest TON OS SE updates we created a fully automated release procedure using Github Actions. In the future this approach could be integrated into TON Labs [tonos-se repository](https://github.com/tonlabs/tonos-se). Details - [Link to CI file](https://github.com/INTONNATION/tonos-se-installers/blob/master/.github/workflows/main.ym), [Link to releases](https://github.com/INTONNATION/tonos-se-installers/releases/).

	



## tonsectl details

tonsectl utility is based on [Cobra](https://github.com/spf13/cobra) library, which is used in many Go projects such as [Kubernetes](http://kubernetes.io/), [Hugo](https://gohugo.io/) and [Github CLI](https://github.com/cli/cli). [This doc](https://github.com/spf13/cobra/blob/master/projects_using_cobra.md) contains a more extensive list of projects using Cobra.  \
In our case Cobra is used in each tonsectl command. Commands are described in correspondent files under [tonsectl/cmd](https://github.com/INTONNATION/tonos-se-installers/tree/master/tonsectl/cmd) directory.

Some of them like _start_, _status_, _stop_ and _reset_ just utilize an API developed under [app/tonseapi/tonseapi.go](https://github.com/INTONNATION/tonos-se-installers/tree/master/tonsectl/app/tonseapi). _init_ command manages an API itself. _install_ - runs init scripts dependent on a GO runtime e.g. Operating System. _start_ will also trigger _init_ first to spin up an API in case it’s not running.

## Prerequirements

TON Q-SERVER which is used in TON OS SE requires Git installed and available in PATH.


## Verification

Navigate to [http://localhost/graphql](http://localhost/graphql) and run the following Graph QL query to verify pre-deployed giver details:


    {
      accounts(
        filter: {
          id: {
            eq: "0:b5e9240fc2d2f1ff8cbb1d1dee7fb7cae155e5f6320e585fcc685698994a19a5"
          }
        }
      ) {
        id
        balance
       code
      }
    }

**NOTE**: our solution was successfully tested using the following test suites:



*   [https://github.com/tonlabs/ton-client-js#run-tests-on-node-js](https://github.com/tonlabs/ton-client-js#run-tests-on-node-js) 
*   [https://github.com/tonlabs/TON-SDK/tree/master/ton_client](https://github.com/tonlabs/TON-SDK/tree/master/ton_client) 


## CLI and SDK configuration


### CLI

     tonos-cli config --url http://localhost

Details - [https://docs.ton.dev/86757ecb2/p/8080e6-tonos-cli](https://docs.ton.dev/86757ecb2/p/8080e6-tonos-cli) 


### SDK

     TonClient.defaultConfig = {
       network: {
        endpoints: ['http://localhost]
       },
     };

Details - [https://docs.ton.dev/86757ecb2/p/5328db-configure-sdk/b/18573c](https://docs.ton.dev/86757ecb2/p/5328db-configure-sdk/b/18573c) 


## Advanced usage

The following table will describe all the abilities of tonsectl and API. 


## tonsectl


<table>
  <tr>
   <td><strong>Command description</strong>
   </td>
   <td><strong>tonsectl</strong>
   </td>
  </tr>
  <tr>
   <td>Install required dependencies and config files into ~/tonse directory
   </td>
   <td><em>tonsectl install</em>
   </td>
  </tr>
  <tr>
   <td>Spin up an API only
   </td>
   <td><em>tonsect init</em>
   </td>
  </tr>
  <tr>
   <td>Spin up an API and start TON OS SE
   </td>
   <td><em>tonsectl start</em>
   </td>
  </tr>
  <tr>
   <td>Stop TON OS SE and API
   </td>
   <td><em>tonsectl stop</em>
   </td>
  </tr>
  <tr>
   <td>Check TON OS SE status
   </td>
   <td><em>tonsect status</em>
   </td>
  </tr>
  <tr>
   <td>Restart TON OS SE from scratch
   </td>
   <td><em>tonsectl reset</em>
   </td>
  </tr>
  <tr>
   <td>Get version
   </td>
   <td><em>tonsectl version</em>
   </td>
  </tr>
  <tr>
   <td>Get help
   </td>
   <td><em>tonsectl help</em>
   </td>
  </tr>
</table>



## API Methods

By default API run on port 10000 and include the following methods:

_Start TONOS SE components_
     
    GET /tonse/start 

_Stop TONOS SE components_
        
    GET /tonse/stop

_Delete TONOS SE components data (require /tonse/stop executed before)_

    GET /tonse/reset

_Return Go Slice [ ] with component PIDs_

    GET /tonse/status




## TONOS SE components custom configuration:

Parameters like ports, data location, loglevel etc can be managed by editing following configuration files:


#### ArangoDB

~/tonse/arangodb/etc/config (%userprofile%/tonse/arangodb/etc/config for Windows)

Details: [https://www.arangodb.com/docs/stable/administration-configuration.html](https://www.arangodb.com/docs/stable/administration-configuration.html) 


#### TON Node

~/tonse/node/cfg (%userprofile%/tonse/node/cfg for Windows) \
~/tonse/node/log_cfg.yml (%userprofile%/tonse/node/log_cfg.yml for Windows)

Details: [https://github.com/tonlabs/ton-labs-node](https://github.com/tonlabs/ton-labs-node) 


#### NGINX

/usr/share/nginx/nginx.conf (%userprofile%/tonse/nginx/nginx.conf for Windows)

Details: [https://www.nginx.com/resources/wiki/start/topics/examples/full/](https://www.nginx.com/resources/wiki/start/topics/examples/full/) 


#### Q Server

~/tonse/graphql/.env (%userprofile%/tonse/graphql/.env)

Details: [https://github.com/tonlabs/ton-q-server#ton-q-server](https://github.com/tonlabs/ton-q-server#ton-q-server) 


## Contacts


#### Telegram



*   @renatSK 
*   @azavodovskyi 
*   @sostrovskyi 

Public group - [https://t.me/intonnationpub](https://t.me/intonnationpub) 

Github - [https://github.com/INTONNATION](https://github.com/INTONNATION) 


## ToDo

Windows:



*   generate configs from templates
*   Rewrite init scripts to use GO

Linux:



*   move Nginx logic to an app to not depend on Linux package manager !!!
*   avoid sudo usage (required for Nginx installation, will be fixed by previous one)
*   generate configs from templates
*   Rewrite init scripts to use GO
*   verify tonsectl on other Linux distributions

OSX:



*   move Nginx logic to an app to not depend on Linux package manager !!!
*   avoid sudo usage (required for Nginx installation, will be fixed by previous one)
*   generate configs from templates
*   Rewrite init scripts to use GO
