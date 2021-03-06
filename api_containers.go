/*
 * Functions API
 *
 * # Introduction  Scaleway Functions is a `Function As A Service` product which gives users the ability to deploy atomic serverless workloads and only pay for resources used while functions are running.  It provides many advantages, such as:  - Functions are only executed when a event is triggered, which allows users to save money while code is not running - Auto-Scalability:   - Automated `Scaling up and down` based on user configuration (e.g. min: 0, max: 100 replicas of my function).   - Automated `Scaling to zero` when function is not executed, which saves some money for the user and save Computing resources for the cloud provider. - Scale only the endpoint  ## Main features  - Fully isolated environments - Scaling to zero (save money and computing resources while code is not executed) - High Availability and Scalability (Automated and configurable, each function may scale automatically according to incoming workloads) - Runtimes for the following programming languages:   - Golang   - node.js v8 and v10   - Python v2.7 and v3.7   - Container As A Service: deploy any non-root container listening on port \\$PORT - Multiple event sources:   - HTTP (request on our Gateway will execute the function)   - CRON (time-based job, runs according to configurable cron schedule) - Integrated with the Scaleway Container Registry product   - Each of your functions namespace has an associated registry namespace   - All your functions are available as docker image in this registry namespace   - Each version of your function matches a tag of this image  # Scaleway Functions Components  ## Namespaces  A `Namespace` is basically a project, a group of `functions`, in which you may set-up `environment variables` to use in each function.  **Please Note** that Scaleway Functions operates upon `Kubernetes`, so we use `Docker` container technology to execute user's code.  Therefore, we integrate our APIs with `Scaleway's Container Registry` product to store user's docker images. **Each Namespace gets a Container Namespace** in which functions images will be pushed.  ## Functions  A `Function` in Scaleway Functions consists of multiple components:  - A `Runtime` (Golang, Python 2/3, Node 8/10...), basically the programming language/environment in which your code will be executed. - `Environment Variables`: You may configure specific environment variables (Database host/credentials for example) which are safely encrypted in our Database, and will be mounted inside your Functions. **Note** that environment variables set at `Namespace` level will also be mounted (in every function). Environment variables written at `function` level override the ones set at `namespace` level (if two env var have the same name for example). - `Source code`: In order to run in the cloud, a function must contain user's source code based on the programming language chosen in `runtime` variable. - `Resources`: users may decide how much computing resources to allocate to each function -> `Memory Limit` (in MB). We will then allocate the right amount of `CPU` based on Memory Limit choice. The right choice for your functions's resources is very important, as you will be billed based on compute usage over time and the number of functions executions.  Representation of given CPU resources based on configured Memory Limit (in MB) for a function:  | Memory (in MB) | CPU  | | :------------: | :--: | |      128       | 70m  | |      256       | 140m | |      512       | 280m | |      1024      | 560m |  Where 560mCPU accounts roughly for half of one CPU power of a Scaleway General Purpose instance  Supported runtimes:  - node8 - node10 - python2 - python3 - golang 1.11+  ## Containers  Containers are applications you deploy with your own runtime:  - Create a docker image - Create a container - Push your image in your registry namespace - Deploy  They are used, scaled and billed like functions  ## CRON  A `CRON` is a type of event which triggers a Scaleway Function (or Container), it is an `add-on` to your function.  CRONs inside Scaleway Serverless have the following properties:  - `schedule`: UNIX Formatted CRON schedule. Your function will be executed based on this schedule. For example, `5 4 * * 0` means execute my function at \"04:05 AM\" on each Sunday (see this [page from Ubuntu's official documentation](https://doc.ubuntu-fr.org/cron)). - `args`: JSON Object passed to your function. You can use this property to define data that will be passed to your function's `event.body` object. For Containers, you might handle these arguments as the HTTP Request's Body.  Under the hood, CRON Triggers are [Kubernetes JOBs](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/) sending HTTP POST requests to your function/container.  To deploy a CRON Function, you must first:  - Create a Namespace - Deploy a function (or a container) - Create a CRON and associate it with the created/deployed function.  ## Authentication  By default, creating a function or a container will make it `public`, meaning that anybody knowing the endpoint could execute it.  A function or a container can be made `private` with the `privacy` parameter.  Here is the workflow used to authenticate to a `private` Scaleway Function:  - Create a function with privacy `private` - Deploy your function - Generate a specific `token` from our API - Send a request to your function and provide the generated token (all unauthenticated requests will be rejected).  ### Tokens  Privacy works with JWT tokens. A JWT Token can be retrieved from the endpoint GET `/jwt/issue`. Depending on the parameters, a jwt token can be valid for either a function, a container, or a namespace:  - `/jwt/issue?namespace_id=1`: issues JWT valid for all functions inside namespace with ID `1`. - `/jwt/issue?function_id=1`: issues JWT valid only for function with id `1`. - `/jwt/issue?container_id=1`: issues JWT valid only for container with id `1`.  **Note that you may (optional) provide an expiration date (formatted \"yyyy-mm-ddT00:00:00Z\") for the token**: example `/jwt/issue?expiration_date=2020-01-02T00:00:00Z&namespace_id=1` will generate a token, valid for all functions and containers inside Namespace with id `1`, and this token will be valid until January 2nd 2020.  The token will have the following claims:  ```json {   \"application_claim\": [     {       \"namespace_id\": \"string\",       \"application_id\": \"string\" // optional: id of function/container     }   ] } ```  Tokens are not stored by Scaleway and can not be retrieved if lost (but new tokens can be generated).  Token revocation is not yet supported, the best way to reset the tokens is to destroy and recreate the namespaces and all of its functions.  ### Functions  A `private` function observes this behaviour:  - If a call is done without `SCW_FUNCTIONS_TOKEN` header, the call is rejected (`Status Code 404`) - If `SCW_FUNCTIONS_TOKEN` header is provided, the token is validated using a public key attached to the namespace.  The environment variables `SCW_PUBLIC`, `SCW_PUBLIC_KEY`, `SCW_NAMESPACE_ID`, `SCW_APPLICATION_ID` are provided by our APIs to validate incoming tokens.  For example, to execute a private function by providing a JWT using `curl`, you may run the following command:  ```bash curl -H \"SCW_FUNCTIONS_TOKEN: <generated-token>\" <your-function-host> ```  ### Containers  As the token validation is done in the function runtime, marking a container as `private` will not do the actual authentication.  Instead, it will set the following environment variables, which you can use in your application to validate incoming requests (token provided by our APIs):  - `SCW_PUBLIC`: `true` or `false` based on your privacy settings - `SCW_PUBLIC_KEY`: PEM-encoded public Key used to decrypt tokens - `SCW_NAMESPACE_ID`: Current Namespace ID - `SCW_APPLICATION_ID`: Current Container ID  As described above, tokens generated from our API will contain either `namespace id` or `application id` in its claims, so you may verify it's validity (after decrypting the JWT with the inject `SCW_PUBLIC_KEY`).  ## Logs  Functions and containers output logs can be retrieved from the endpoint GET `/logs`. You need to pass its ID as an `application_id` parameter.  # Quick Start Guide  ## Pre-requisites  Whether you decide to use Serverless Framework or directly our API, you'll need your Scaleway Organization ID and a Scaleway Organization Access Key.  - Install `curl` - Install `jq` will make it easier to manage JSON output from our APIs  To call Scaleway API, you need an `X-Auth-Token`. If you don't have one yet, you can create it on the [credentials page](https://cloud.scaleway.com/#/credentials) of your Scaleway account (must be done via web interface).  In order to retrieve your `Organization ID` and your `secret Key`, you must go to your [console's credentials page](https://console.scaleway.com/account/credentials):  - Login/Register to [Scaleway console](https://console.scaleway.com) - Go to your [credentials management page](https://console.scaleway.com/account/credentials) - Retrieve your `organization ID` and generate a token (see following picture):   ![credentials section](https://functions-doc.s3.fr-par.scw.cloud/credentials_section.png) - Retrieve your token's `secret key`:   ![token secret key](https://functions-doc.s3.fr-par.scw.cloud/secret_key.png)  Then, export then as variables to use them with curl  ```bash export TOKEN=\"<Secret key of your token>\" # Only available in fr-par at the moment export REGION=\"<choose your location (nl-ams/fr-par)>\" export ORGANIZATION_ID=\"<your organization ID>\" ```  ## Serverless Framework  The following sections explain how to use our API, with a tutorial and the auto-generated API documentation. However, we developed a [Serverless Framework plugin](https://github.com/scaleway/serverless-scaleway-functions) enabling users to deploy their serverless workloads much more easily with a single `serverless deploy` command. No magic there, it's just a nice tool calling our API.  If what you are looking for is an easy way to deploy your code, you may prefer Serverless Framework.  Below, you will find a step-by-step guide on how to create a `namespace`, configure and deploy `functions`, and trigger your `functions` via HTTP and CRON.  ## Create a Namespace  Customize the name and set your organization ID  ```bash  curl -X POST \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/namespaces\" -H \"accept: application/json\" -H \"X-Auth-Token: $TOKEN\" -H \"Content-Type: application/json\" \\ -d \"{\\\"name\\\": \\\"your-namespace-name\\\", \\\"organization_id\\\": \\\"$ORGANIZATION_ID\\\", \\\"environment_variables\\\": {\\\"YOUR_VARIABLE\\\": \\\"content\\\"}}\" ```  Copy the `id` field of the response to use at the next steps. For the sake of simplicity we will save the ID to a variable, which we will use in the following examples:  ```bash export NAMESPACE_ID=\"<your namespace id>\" ```  To destroy a namespace (along with all functions and crons) use the following call:  ```bash curl -s -H \"X-Auth-Token: $TOKEN\" -X DELETE \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/namespaces/$NAMESPACE_ID\" ```  ## Write a function handler  **Please note that our runtimes are AWS Lambda Compatible**, which means that we respect Lambda's format (event, context and callback parameters hold the same keys as AWS Lambda's). Be careful about `context` though, as we do not provide the exact same keys as Lambda (we don't have cognito services for example).  For this example, we'll be using `node10` runtime:  ```bash touch handler.js ```  And inside `handler.js` file:  ```javascript // handler.js module.exports.myHandler = async (event, context, callback) => {   const response = {     body: JSON.stringify({message: \"Hello, World\"}),     statusCode: 200,     headers: {       MY_HEADER: \"its content\",     },   }   return response } ```  ## Create a function  When creating a function, you may customize multiple fields:  - `name`: The name of your function - `namespace_id`: ID of the namespace in which you want to create your function - `runtime`: Your function's runtime, check the supported runtimes above - `memory_limit`: Memory (in MB) allocated to your function, see the table of memory/CPU allocation above (increasing the memory limit will increase the cost of your function executions as we allocate more resources to your functions). - `min_scale`: Minimum replicas for your function, defaults to `0`, **Note** that a function is `billed` when it gets executed, and using a `min_scale` greater than 0 will cause your function to run all the time. - `max_scale`: Maximum replicas for your function (defaults to `20`), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured `max_scale`. - `handler` (More details with examples in each language/runtime section below):   - `Python`: Path to function handler's file and the function to use as the handler: `src/handler.my_handler` => file `handler.py` defining a `my_handler` function, inside `src` folder.   - `Node`: Path to function handler's file, suffixed the name of the function to use as the handler: `src/handler.myHandler` => file `handler.js` exporting a `myHandler` function, inside `src` directory.   - `Golang`: Path to the package containing the handler: `my_handler`: the code containing the handler is located inside a `my_handler` directory (must be `package main`, and exposing a `main function`).  ```bash curl -X POST \\ -H \"X-Auth-Token: $TOKEN\"\\ \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/functions\"\\ -d \"{\\\"name\\\": \\\"function-name\\\", \\\"namespace_id\\\": \\\"$NAMESPACE_ID\\\", \\\"memory_limit\\\": 128, \\\"min_scale\\\": 0, \\\"max_scale\\\": 20, \\\"runtime\\\": \\\"node10\\\", \\\"handler\\\": \\\"handler.myHandler\\\"}\" export FUNCTION_ID = \"<your-function-id>\" ```  ## Upload your source code (Scaleway runtime)  These steps only apply if you use a Scaleway runtime. In this section, you will upload your code to a S3 bucket, which we'll package and build into a container image.  This container image will then be available in a registry namespace associated to your functions namespace.  ### Archive your code  ```bash export FUNCTION_ARCHIVE=\"function-$FUNCTION_ID.zip\" ```  You may then create a zip archive with your code:  ```bash zip $FUNCTION_ARCHIVE handler.js ```  **Please Note that if you wish to use external dependencies, you will have to package them inside the zip archive as well**:  ```bash zip -r $FUNCTION_ARCHIVE package.json handler.js node_modules ```  ### Get a presigned URL for our S3 Bucket to store your function handler  You need to get the size of your archive in bytes, in order to ask for a presigned URL to upload your source code:  ```bash ls -lh -rw-r--r--  1 user  group   675 Apr 18 15:42 $FUNCTION_ARCHIVE  export ARCHIVE_SIZE=675 ```  ```bash curl -X GET -H \"X-Auth-Token: ${TOKEN}\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/functions/$FUNCTION_ID/upload-url?content_length=$ARCHIVE_SIZE\"  # Example of response from our API {\"url\":\"https://s3.fr-par.scw.cloud/scw-database-srvless-prod/uploads/function-b0525a73-947d-4ba4-92de-17f267a7ec5a.zip?X-Amz-Algorithm=AWS4-HMAC-SHA256\\u0026X-Amz-Credential=SCW6Z6VKJVG81FQZVB14%2F20190627%2Ffr-par%2Fs3%2Faws4_request\\u0026X-Amz-Date=20190627T092839Z\\u0026X-Amz-Expires=3600\\u0026X-Amz-SignedHeaders=content-length%3Bcontent-type%3Bhost\\u0026X-Amz-Signature=e9f3e22f39638dac047f0f4e9ab521c7971cacb01f61f523cb948baa328a0eff\",\"headers\":{\"content-length\":[\"347\"],\"content-type\":[\"application/octet-stream\"]}} ```  As you can see, the url is not properly formatted (\\u0026...), in order to use it properly to upload your code, you must copy the full URL with quotes (otherwise your terminal might add unwanted \\ in the url string):  ```bash export FUNCTION_ARCHIVE_URL=$(echo -n \"<your-url>\") ```  **Note**: that you will get an error in the following step if you do not Copy the url with wrapping quotes and save inside a variable using echo -n to replace \\u0026 expressions  If you use postman, you can usually export the presigned url as it is, as long as you copy/paste the quotes too.  ### Upload your code to the presigned URL  ```bash curl -H \"Content-Type: application/octet-stream\" --upload-file $FUNCTION_ARCHIVE -H \"Content-Length: $ARCHIVE_SIZE\" $FUNCTION_ARCHIVE_URL ```  ## Deploy a function  Then, run the following command to deploy your function:  ```bash curl -X POST -H \"X-Auth-Token: $TOKEN\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/functions/$FUNCTION_ID/deploy\" -d \"{}\" ```  The process may take a little bit of time, as we have to build your source code into an executable function (wrapped by our runtimes), and deploy it to our cloud platform.  ## Trigger your function  Once your function has been properly deployed, you may retrieve your function's HTTP(s) endpoint with the following command:  ```bash curl -X GET -H \"X-Auth-Token: $TOKEN\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/functions/$FUNCTION_ID\"  export FUNCTION_ENDPOINT=\"<endpoint>\" ```  And then, call your function via its endpoint:  ```bash curl -X GET \"$FUNCTION_ENDPOINT\" ```  ## Get your functions logs  To retrieve the functions output logs:  ```bash curl -X GET -H \"X-Auth-Token: $TOKEN\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/logs?application_id=$FUNCTION_ID\" ```  ## Create a CRON Trigger for your Function  As described above, `CRON` triggers are a way to execute your applications (Functions and Containers) periodically, based on a given Schedule.  It means that we can execute our function every day at 1PM for example, with a given set of data.  In order to add a `CRON` Trigger to your function, you need to retrieve your function ID (Done previously if you followed the guide), and create a new CRON associated to your function:  ```bash curl -X POST -H \"X-Auth-Token: $TOKEN\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/crons\" -d \"{\\\"application_id\\\": \\\"$FUNCTION_ID\\\", \\\"schedule\\\": \\\"0 13 * * *\\\", \\\"args\\\": {\\\"key\\\": \\\"value\\\"}}\" ```  The above request will create and deploy a Kubernetes CRON Job in charge of executing your function every day at 13:00, with the data `{\"key\": \"value\"}`, retrieved from the `event.body` object in your handler.  **Note that this step is also applicable to containers, you just need to pass your container ID as the CRON's `application_id` property**.  ## Create a container  Creating a container is nearly identical to creating a function.  The main difference is that here you don't need to upload your code in a S3 bucket. Instead, you need to build it as a docker image and push it to our registry.  ```bash curl -X POST \\ -H \"X-Auth-Token: $TOKEN\" \\ \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/containers\" \\ -d \"{\\\"name\\\": \\\"container-name\\\", \\\"namespace_id\\\": \\\"$NAMESPACE_ID\\\", \\\"memory_limit\\\": 128, \\\"min_scale\\\": 0, \\\"max_scale\\\": 20}\" ```  Let's export the container_name for later:  ```bash export CONTAINER_NAME=\"<container_name>\" export CONTAINER_ID=\"<container_id>\" ```  ### Get your registry namespace  First, get your registry_namespace_id  ```bash curl -X GET -H \"X-Auth-Token: ${TOKEN}\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/namespaces/$NAMESPACE_ID\" export REGISTRY_NAMESPACE_ID=\"<registry-namespace-id>\" ```  Secondly, get your registry name  ```bash curl -X GET -H \"X-Auth-Token: ${TOKEN}\" \"https://api.scaleway.com/registry/v1beta2/regions/$REGION/namespaces/$REGISTRY_NAMESPACE_ID\" export REGISTRY_ENDPOINT=\"<endpoint>\" ```  ### Push your image  We suppose you already have a working image here. It can be anything which listens on a env variable \\$PORT variable. Note that we run your container as user 1000, not root, so it must be runnable under these conditions  To push your image, we invite you to check the container registry documentation.  ```bash docker login $REGISTRY_ENDPOINT -u userdoesnotmatter -p $TOKEN docker tag myimage $REGISTRY_ENDPOINT/${CONTAINER_NAME}:latest docker push $REGISTRY_ENDPOINT/${CONTAINER_NAME}:latest ```  ### Deploy a container  ```bash curl -X POST -H \"X-Auth-Token: $TOKEN\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/containers/$CONTAINER_ID/deploy\" -d \"{}\" ```  ## Trigger your container  ```bash curl -X GET -H \"X-Auth-Token: $TOKEN\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/containers/$CONTAINER_ID\" export CONTAINER_ENDPOINT=\"<endpoint>\"  curl -X GET \"$CONTAINER_ENDPOINT\" ```  ## Get your containers logs  To retrieve the containers output logs:  ```bash curl -X GET -H \"X-Auth-Token: $TOKEN\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/logs?application_id=$CONTAINER_ID\" ```  # Writing code  Runtimes are environment that you may use to develop their cloud functions.  Scaleway Runtimes are `Lambda Compatible` -> For API Gateway Proxy `event types` (as we only support HTTP and CRON, but cron basically sends HTTP requests to deployed functions).  In a common use case with `Serverless Framework` for example, in which a user has multiple functions in the same repository and would like to upload them all at the same time with a single command, we need a way to know, which file (Python/JavaScript) or package (Golang) to use to execute our functions.  **Please Note** that in some runtimes, this `decision` is made at runtime (Python and JavaScript, as they are interpreted languages, so our runtime will only `import` the handler), while in some others (Golang), it is done at compile/build time (when user `deploys` a function) as we need to build user's package.  ## Node (v8, v10)  ### Node handler function  **Please Note** currently function handler must be a named exported component => `module.exports.myHandler = (event, context, callback) => {}`  There are multiple ways to return a response from a handler function:  First one: `return object with body and statusCode` will set the status code as HTTP Response Code, and body as the Response's body, headers as Headers.  - Stringified **body** (like `AWS Lambda`):  ```javascript module.exports.myHandler = (event, context, callback) => {   return {     statusCode: 201,     body: JSON.stringify({       message: \"async function\",     }),     headers: {       \"Content-Type\": \"application/json\",     },   } } ```  - **Not** Stringified **body** (like `AWS Lambda`):  ```javascript module.exports.myHandler = (event, context, callback) => {   return {     statusCode: 201,     body: {       message: \"async function\",     },     headers: {       \"Content-Type\": \"application/json\",     },   } } ```  Second: `return Object/entity (number, boolean, string...) withtout properties body and statusCode` will return the response `as is`:  ```javascript module.exports.myHandler = (event, context, callback) => {   return {     message: \"message\",   }    // Or   return JSON.stringify({message: \"message\"})   // OR   return \"Hello, world\"   // OR   return 1 // true,false,undefined,null... } ```  `Use Callback parameter`:  ```javascript module.exports.myHandler = (event, context, callback) => {   const response = {     statusCode: 201,     body: {       message: \"async function\",     },     headers: {       \"Content-Type\": \"application/json\",     },   }    // Successful response   callback(undefined, response)   // Error response   callback(err) } ```  **Note that you may use life changing async in handlers** :)  `return a Promise`:  ```javascript module.exports.myHandler = async (event, context, callback) => {   return {     statusCode: 201,     body: {       message: \"async function\",     },     headers: {       \"Content-Type\": \"application/json\",     },   } }  // OR module.exports.myHandler = (event, context, callback) => {   const response = {     statusCode: 201,     body: {       message: \"async function\",     },     headers: {       \"Content-Type\": \"application/json\",     },   }    return new Promise((resolve, reject) => {     // do something     if (err) return reject(err)     return resolve(response)   }) } ```  ### Node handler name  The Handler name is basically a path to the handler file.  For example, let's say I have two handlers `hello.js` and `world.js` inside `src/handlers` folder:  ``` src -- handlers ---- hello.js => module.exports.sayHello ---- world.js => module.exports.toTheWorld ```  Then, you need to provide a custom handler name for each of these handlers, so each of your functions will know which handler file to run: `hello` -> `src/handlers/hello.sayHello` and `world` -> `src/handlers/world.toTheWorld`.  **By default**, the handler path is `handler.handle` (`module.exports.handle` in handler.js).  ### Node additional dependencies  If you ever need to push external dependencies for your node.js functions, you will have to package your `node_modules` directory into your deployment archive.  ``` -- handler.js -- node_modules ---- <your-dependencies> ```  You may definitely use tools such as `webpack` or [NCC](https://github.com/zeit/ncc) (CLI tool to build node.js executables, inspired from `go` CLI), which will package your code into separate files, you will then be able to upload your compiled handler file, which reduces the size of your bundle.  Example:  ```bash ncc handler.js -o build/handler.js # -> Builds dist/inde ```  Then, set up your `function handler` to be: `build/handler.js` if you package the whole `build` directory. Don't forget to point the `function handler` property to the path of your built handler in your archive (if `build/handler.myHandler` then `handler must be build/handler.js`)  ## Golang  **Only versions 1.11+ are supported on Scaleway Serverless**  ### Golang handler function  **Please Note** that every handler must be in its package, identified by `package main`, and exporting a main function with the following `lambda.Start` statement:  ```golang // Must Always be package main package main  // Import both packages events and lambda from scaleway-functions-go library import (  \"encoding/json\"  \"github.com/scaleway/scaleway-functions-go/events\"  \"github.com/scaleway/scaleway-functions-go/lambda\" )  // Handler - Your handler function, uses APIGatewayProxy event type as your function will always get HTTP formatted events, even for CRON func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {  return events.APIGatewayProxyResponse{   Body:       \"Your response\",   StatusCode: 200,  }, nil }  // Main function is mandatory -> Must call lambda.Start(yourHandler) otherwhise your handler will not be called properly. func main() {  lambda.Start(Handler) } ```  ### Golang handler name  In Golang, as it is a `compiled` language, you need to provide Scaleway Function API with a `handler name` pointing to the function's directory.  ``` src -- handlers ---- hello ------ go.mod ------ go.sum ------ main.go -> package main in \"handlers/hello\" subdirectory ---- world ------ go.mod ------ go.sum ------ main.go -> package main in \"handlers/workd\" subdirectory handler.go -> package main at the root of serverless project go.mod go.sum ```  Then, you need to provide a custom handler name for each of these handlers, so each of your function will know which handler file to run: `hello` -> `handlers/hello`, `world` -> `handlers/world`, and for the `handler.go` at the root of the project -> `.`  **By default**, the handler path is `.` (`package main` at the root of the archive).  ### Golang additional dependencies  If you need external dependencies for your `Golang handlers`, you may provide these dependencies by using `Go Modules`:  - Our runtimes automatically installs your dependencies at Build time (Once you start the function `deployment`). **Note that dependencies installation at build-time will result in longer builds**.    ```   -- handler.go   -- go.mod   -- go.sum   ```  - You may package your dependencies with the command `go mod vendor`, and provide your generated `vendor` directory to the function archive. **This approach will save you some time during builds**:   ```   -- handler.go   -- go.mod   -- go.sum   -- vendor # Dependencies should be installed inside your vendor directory   ---- <your-dependencies>   ```  ## Python  ### Python handler function  There are multiple ways to return a response from a handler function:  - Classical response object with `HTTP` informations:  ```python def my_handler(event, context):     return {         \"body\": {             \"test\": \"test\"         },         \"statusCode\": 200,         \"headers\": {             \"your-header\": \"your-value\"         }     } ```  - Straight response without `body`:  ```python def my_handler(event, context):     return {\"message\": \"whatever\"}     # or     return \"my Message\" ```  - Stringified response **body** (`AWS Lambda`):  ```python import json  def my_handler(event, context):   return {     \"body\": json.dumps({\"message\": \"Hello\"}),     \"statusCode\": 200,   } ```  ### Python handler name  The Handler name is basically a path to the handler file, suffixed with the function name to use as a handler.  For example, let's say you have two handlers `hello.py` and `world.py` inside `src/handlers` folder:  ```  src -- handlers ---- hello.py => def say_hello ---- world.py => def to_the_world ```  Then, you need to provide a custom handler name for each of these handlers, so each of your functions will know which handler file to run: `hello` -> `src/handlers/hello.say_hello` and `world` -> `src/handlers/world.to_the_world`.  **By default**, the handler path is `handler.handle` (def `handle` in handler.py).  ### Python additional dependencies  Additional dependencies must be included inside a `package` directory at the root of your archive/project:  ```bash # At the root of your archive mkdir package ```  ``` - requirements.txt - handlers --- handler.py => import requests --- secondHandler.py => import requests - package --- requests --- ... ```  #### Standard dependencies  You may install your dependencies to the `package` directory:  ```bash pip install requests --target ./package ```  Or with a `requirements.txt` file:  ```bash pip install -r requirements.txt --target ./package ```  #### Specific libraries (with needs for specific C compiled code)  In some cases, you might need to install libraries which require specific C compiled code such as (for example):  - `numpy` - `tensorflow` - `pandas` - `scikit-learn` - ...  Our Python runtimes run on top of `alpine linux` environments, for these specifics dependencies, you will have to install your dependencies inside a `docker container`, with a specific image, that we are providing to our users.  You may run the following command from the root of your project to install your dependencies before uploading your source code and deploying your function:  ```bash docker run --rm -v $(pwd):/home/app/function --workdir /home/app/function rg.fr-par.scw.cloud/scwfunctionsruntimes/python-dep3:v4.0.0 pip install -r requirements.txt --target ./package ```  This command will run `pip install` with given `requirements.txt` file inside a `docker container` compatible with our function runtimes, and pull the installed dependencies locally to your `package` directory. As these dependencies have been installed on top of `alpine linux` with our compatible `system libraries`, you will be able to upload your source code and deploy your function properly.  **Note** that the example below uses `python3` runtime, but you can easily change the docker image from `rg.fr-par.scw.cloud/scwfunctionsruntimes/python-dep3:v4.0.0` to `rg.fr-par.scw.cloud/scwfunctionsruntimes/python-dep2:v4.0.0`  ## Remove a Scaleway Functions namespace (Project)  When deleting a Functions Namespace, we take care of removing all sub-resources such as `Functions` and `CRONs` deployed in this namespace.  - With `Serverless Framework`:  ```bash serverless remove ```  - With `curl`:  ```bash curl -X DELETE -H \"X-Auth-Token: ${TOKEN}\" \"https://api.scaleway.com/functions/v1alpha2/regions/$REGION/namespaces/$NAMESPACE_ID\" ```  **Please note** that deleting a `Scaleway Functions` namespace will not automatically delete Scaleway Container Registry namespaces linked to your FAAS project. It is your responsibility to manually remove your Registry namespaces via Scaleway Console or API.  ## Develop Locally  If you are using either `node` or `python` runtimes, you may use our [Offline Gateway plugin](https://github.com/scaleway/serverless-offline-scaleway) (Developed and maintained by Scaleway).  ```bash npm install --save-dev serverless-offline-scaleway ```  And in your serverless File:  ```yml plugins:   - serverless-scaleway-functions   - serverless-offline-scaleway ```  You may now invoke your functions locally via the following command line:  ```bash serverless offline start ```  ## API Reference  ### Event  #### Node/Python  - `pathParameters`: map(string)string - Parameters defined in the path of the HTTP Request - `queryStringParameters`: map(string)string - Query Strings parameters of the HTTP Request - `body`: string|byte() - Body of the HTTP Request, you will have to parse it in your handler to use it properly. - `headers`: map(string)string - HTTP Request Headers - `method`: string - HTTP method used - `isBase64Encoded`: boolean - Whether the request body is base64 encoded.  #### Golang  You may take a look at [our scaleway-functions-go package](https://github.com/scaleway/scaleway-functions-go) (events and lambda) packages.  ### Context  Context typings will be supported soon, you may already use it in functions, but every value be `defaults values` as our API does not support it right now. 
 *
 * API version: v1alpha2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ContainersApiService ContainersApi service
type ContainersApiService service

/*
CreateContainer Create a new container
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region The region you want to target
 * @param inlineObject
@return ScalewayFunctionsV1alpha2Container
*/
func (a *ContainersApiService) CreateContainer(ctx _context.Context, region string, inlineObject InlineObject) (ScalewayFunctionsV1alpha2Container, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScalewayFunctionsV1alpha2Container
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/functions/v1alpha2/regions/{region}/containers"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", _neturl.QueryEscape(parameterToString(region, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &inlineObject
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Auth-Token"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
DeleteContainer Delete a container
Delete the container associated with the given id.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region The region you want to target
 * @param containerId
@return ScalewayFunctionsV1alpha2Container
*/
func (a *ContainersApiService) DeleteContainer(ctx _context.Context, region string, containerId string) (ScalewayFunctionsV1alpha2Container, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScalewayFunctionsV1alpha2Container
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/functions/v1alpha2/regions/{region}/containers/{container_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", _neturl.QueryEscape(parameterToString(region, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"container_id"+"}", _neturl.QueryEscape(parameterToString(containerId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Auth-Token"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
GetContainer Get a container
Get the container associated with the given id.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region The region you want to target
 * @param containerId
@return ScalewayFunctionsV1alpha2Container
*/
func (a *ContainersApiService) GetContainer(ctx _context.Context, region string, containerId string) (ScalewayFunctionsV1alpha2Container, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScalewayFunctionsV1alpha2Container
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/functions/v1alpha2/regions/{region}/containers/{container_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", _neturl.QueryEscape(parameterToString(region, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"container_id"+"}", _neturl.QueryEscape(parameterToString(containerId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Auth-Token"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListContainersOpts Optional parameters for the method 'ListContainers'
type ListContainersOpts struct {
    Page optional.Float32
    PageSize optional.Float32
    OrderBy optional.Interface
    NamespaceId optional.String
    Name optional.String
    OrganizationId optional.String
}

/*
ListContainers List all your containers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region The region you want to target
 * @param optional nil or *ListContainersOpts - Optional Parameters:
 * @param "Page" (optional.Float32) -  Page number
 * @param "PageSize" (optional.Float32) -  Page size
 * @param "OrderBy" (optional.Interface of ScalewayFunctionsV1alpha2ListContainersRequestOrderBy) - 
 * @param "NamespaceId" (optional.String) - 
 * @param "Name" (optional.String) - 
 * @param "OrganizationId" (optional.String) - 
@return ScalewayFunctionsV1alpha2ListContainersResponse
*/
func (a *ContainersApiService) ListContainers(ctx _context.Context, region string, localVarOptionals *ListContainersOpts) (ScalewayFunctionsV1alpha2ListContainersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScalewayFunctionsV1alpha2ListContainersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/functions/v1alpha2/regions/{region}/containers"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", _neturl.QueryEscape(parameterToString(region, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderBy.IsSet() {
		localVarQueryParams.Add("order_by", parameterToString(localVarOptionals.OrderBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NamespaceId.IsSet() {
		localVarQueryParams.Add("namespace_id", parameterToString(localVarOptionals.NamespaceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrganizationId.IsSet() {
		localVarQueryParams.Add("organization_id", parameterToString(localVarOptionals.OrganizationId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Auth-Token"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
UpdateContainer Update an existing container
Update the container associated with the given id.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region The region you want to target
 * @param containerId
 * @param inlineObject1
@return ScalewayFunctionsV1alpha2Container
*/
func (a *ContainersApiService) UpdateContainer(ctx _context.Context, region string, containerId string, inlineObject1 InlineObject1) (ScalewayFunctionsV1alpha2Container, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ScalewayFunctionsV1alpha2Container
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/functions/v1alpha2/regions/{region}/containers/{container_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", _neturl.QueryEscape(parameterToString(region, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"container_id"+"}", _neturl.QueryEscape(parameterToString(containerId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &inlineObject1
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Auth-Token"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
