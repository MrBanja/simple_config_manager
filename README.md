# Simple Config Manager: Simplifying Configuration Management

The Simple Config Manager is an essential tool for smaller projects with simple CI/CD, as it allows you to store your configuration in a template file and populate it with environment variables from CI. This tool removes the need for using `sed` or `awk` in your CI pipelines, and eliminates the necessity of `Ansible` or similar tools to manage your configuration files.

## Getting Started
To utilize the Simple Config Manager, follow the steps below:

### Syntax and Parameters
The following syntax and parameters are used for the Simple Config Manager command:
```
Usage of ./scmanager:
  --env_prefix string
        Prefix for the environment variable
  --error_on_empty
        Returns an error if the environment variable value is empty
  --ilepath string
        Path to the configuration template file
  --utput_filepath string
        Path to the output configuration file
```

### Create Template
Start by creating a template configuration file, `conf.tmpl.yaml`, which can be conveniently stored within your repository.
```yaml
db:
  user: {{ index . "DB_USER" }}
  password: {{ index . "DB_PASSWORD" }}
  port: 5643
```
Please note that the variables must be in uppercase letters. This is **crucial** for the template to work correctly.

### Define Environment Variables
Next, define and export your environment variables. These can be stored in CI secrets, for example:
```sh
export WEBAPI_DB_USER="service"
export WEBAPI_DB_PASSWORD="psswd"
```

Environment variables and prefix are case-insensitive, so `WEBAPI_DB_USER` and `webApI_dB_uSeR` are the same.

### Execute Command
Run the Simple Config Manager command as follows:
```sh
/scmanager --filepath=/path/to/conf.tmpl.yaml \
  --env_prefix=WEBAPI_ \
  --output_filepath=/path/to/conf.yaml \
  --error_on_empty
```

### Output
The command will generate (or **override**, if it already exists) the configuration file at `/path/to/conf.yaml` with the following structure:
```yaml
db:
  user: service
  password: psswd
  port: 5643
```

This file will now be filled with your previously defined environment variables, streamlining the configuration management for your project. The Simple Config Manager provides an effective, simple, and secure way to manage your project's configuration.