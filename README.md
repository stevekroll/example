## Environment
#### Create a new `.env` in the root directory of the project.

Here is a basic example:
```env
APP=example
VERSION=0.1.0
HOST=http://127.0.0.1
PORT=8080
LOG_LEVEL=-4
TIMEOUT=5000
```

## Taskfile

Installation:
https://taskfile.dev/installation/

We will use this to chain together multiple commands for simplifying processes with more than one step.

For example, running the webserver: 
```
task webserver
```
vs
```
docker compose up
```
```
go run main.go webserver
```
```
docker compose down
```

## K6
#### Build Custom K6 Image
* Run the command `task k6:pull` to pull the latest xk6 image.
* Then run `task k6:build` to build a custom k6 image with faker extension.

#### Running A Load Test
* Run `task k6:new filename={$filename}` to create a new load test script.
* Run the script with command `task k6:run filename={$filename}`.

## Grafana
#### Configure Grafana Dashboard
* Open `http://localhost:3000`.
* Select 'Create New Dashboard'.
* Create a new data source and select `Prometheus`.
* Use http://host.docker.internal:9090.
* Import a new dashboard using JSON found in `tools/grafana/dashboards`.

## Prometheus
#### Run Unit Tests
* Make sure both docker compose and prometheus are running.
* Run the command `task prometheus:test`.

## PPROF

```
usage:

Produce output in the specified format.

   pprof <format> [options] [binary] <source> ...

Omit the format to get an interactive shell whose commands can be used
to generate various views of a profile

   pprof [options] [binary] <source> ...

Omit the format and provide the "-http" flag to get an interactive web
interface at the specified host:port that can be used to navigate through
various views of a profile.

   pprof -http [host]:[port] [options] [binary] <source> ...

Details:
  Output formats (select at most one):
    -callgrind       Outputs a graph in callgrind format
    -comments        Output all profile comments
    -disasm          Output assembly listings annotated with samples
    -dot             Outputs a graph in DOT format
    -eog             Visualize graph through eog
    -evince          Visualize graph through evince
    -gif             Outputs a graph image in GIF format
    -gv              Visualize graph through gv
    -kcachegrind     Visualize report in KCachegrind
    -list            Output annotated source for functions matching regexp
    -pdf             Outputs a graph in PDF format
    -peek            Output callers/callees of functions matching regexp
    -png             Outputs a graph image in PNG format
    -proto           Outputs the profile in compressed protobuf format
    -ps              Outputs a graph in PS format
    -raw             Outputs a text representation of the raw profile
    -svg             Outputs a graph in SVG format
    -tags            Outputs all tags in the profile
    -text            Outputs top entries in text form
    -top             Outputs top entries in text form
    -topproto        Outputs top entries in compressed protobuf format
    -traces          Outputs all profile samples in text form
    -tree            Outputs a text rendering of call graph
    -web             Visualize graph through web browser
    -weblist         Display annotated source in a web browser

  Options:
    -call_tree       Create a context-sensitive call tree
    -compact_labels  Show minimal headers
    -divide_by       Ratio to divide all samples before visualization
    -drop_negative   Ignore negative differences
    -edgefraction    Hide edges below <f>*total
    -focus           Restricts to samples going through a node matching regexp
    -hide            Skips nodes matching regexp
    -ignore          Skips paths going through any nodes matching regexp
    -intel_syntax    Show assembly in Intel syntax
    -mean            Average sample value over first value (count)
    -nodecount       Max number of nodes to show
    -nodefraction    Hide nodes below <f>*total
    -noinlines       Ignore inlines.
    -normalize       Scales profile based on the base profile.
    -output          Output filename for file-based outputs
    -prune_from      Drops any functions below the matched frame.
    -relative_percentages Show percentages relative to focused subgraph
    -sample_index    Sample value to report (0-based index or name)
    -show            Only show nodes matching regexp
    -show_from       Drops functions above the highest matched frame.
    -source_path     Search path for source files
    -tagfocus        Restricts to samples with tags in range or matched by regexp
    -taghide         Skip tags matching this regexp
    -tagignore       Discard samples with tags in range or matched by regexp
    -tagleaf         Adds pseudo stack frames for labels key/value pairs at the callstack leaf.
    -tagroot         Adds pseudo stack frames for labels key/value pairs at the callstack root.
    -tagshow         Only consider tags matching this regexp
    -trim            Honor nodefraction/edgefraction/nodecount defaults
    -trim_path       Path to trim from source paths before search
    -unit            Measurement units to display

  Option groups (only set one per group):
    granularity      
      -functions       Aggregate at the function level.
      -filefunctions   Aggregate at the function level.
      -files           Aggregate at the file level.
      -lines           Aggregate at the source code line level.
      -addresses       Aggregate at the address level.
    sort             
      -cum             Sort entries based on cumulative weight
      -flat            Sort entries based on own weight

  Source options:
    -seconds              Duration for time-based profile collection
    -timeout              Timeout in seconds for profile collection
    -buildid              Override build id for main binary
    -add_comment          Free-form annotation to add to the profile
                          Displayed on some reports or with pprof -comments
    -diff_base source     Source of base profile for comparison
    -base source          Source of base profile for profile subtraction
    profile.pb.gz         Profile in compressed protobuf format
    legacy_profile        Profile in legacy pprof format
    http://host/profile   URL for profile handler to retrieve
    -symbolize=           Controls source of symbol information
      none                  Do not attempt symbolization
      local                 Examine only local binaries
      fastlocal             Only get function names from local binaries
      remote                Do not examine local binaries
      force                 Force re-symbolization
    Binary                  Local path or build id of binary for symbolization
    -tls_cert             TLS client certificate file for fetching profile and symbols
    -tls_key              TLS private key file for fetching profile and symbols
    -tls_ca               TLS CA certs file for fetching profile and symbols

  Misc options:
   -http              Provide web interface at host:port.
                      Host is optional and 'localhost' by default.
                      Port is optional and a randomly available port by default.
   -no_browser        Skip opening a browser for the interactive web UI.
   -tools             Search path for object tools

  Legacy convenience options:
   -inuse_space           Same as -sample_index=inuse_space
   -inuse_objects         Same as -sample_index=inuse_objects
   -alloc_space           Same as -sample_index=alloc_space
   -alloc_objects         Same as -sample_index=alloc_objects
   -total_delay           Same as -sample_index=delay
   -contentions           Same as -sample_index=contentions
   -mean_delay            Same as -mean -sample_index=delay

  Environment Variables:
   PPROF_TMPDIR       Location for saved profiles (default $HOME/pprof)
   PPROF_TOOLS        Search path for object-level tools
   PPROF_BINARY_PATH  Search path for local binary files
                      default: $HOME/pprof/binaries
                      searches $buildid/$name, $buildid/*, $path/$buildid,
                      ${buildid:0:2}/${buildid:2}.debug, $name, $path,
                      ${name}.debug, $dir/.debug/${name}.debug,
                      usr/lib/debug/$dir/${name}.debug
   * On Windows, %USERPROFILE% is used instead of $HOME
```