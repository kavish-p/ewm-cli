# CLI program for IBM Enterprise Workflow Management (EWM)
CLI program to interact with IBM EWM for tasks such as resolving work items and creating new defect workitems.
IBM EWM was formerly named Rational Team Concert (RTC)

Build Executable Binary
============
```
go build
```

Configuration
============
The config file is located at $HOME/.ewm-cli.yaml by default and should be created before using the CLI commands. The oslc_context, filedAgainstCategory, and defectType attributes are not required for `get context` and `get category` commands.

Config files can also be passed through the `--config` flag.

Below is a sample configuration file.
```
base_url: https://10.168.0.61:9443
ewm_username: kavish
ewm_password: P@ssw0rd
oslc_context: _isel8HoTEeyoCYEXLNJaZA
filedAgainstCategory: _l-ZCoHoTEeyoCYEXLNJaZA
defectType: com.ibm.team.workitem.workItemType.defect
```

Usage Examples
============

### Find the project/context IDs on EWM server
```
ewm-cli get context
```

### Get FiledAgainst Category IDs for a specific project
```
ewm-cli get category --context _isel8HoTEeyoCYEXLNJaZA
```

### Create a defect on IBM EWM
```
ewm-cli create defect --description "this is a description" --summary "Unable to call API endpoint /abc"
```

### Check Postman/Newman JSON output file and create a defect workitem on IBM EWM for each failed test case.
```
ewm-cli check newman --report C:\Users\Kavish\Documents\Temp\test\report.json
```

### Check Pytest JSON output file and create a defect workitem on IBM EWM for each failed test case.
The Pytest JSON Report is generated by the [pytest-json-report](https://pypi.org/project/pytest-json-report/) pip package.
```
ewm-cli check pytest --report C:\Users\Kavish\Documents\Temp\test\report.json
```

### Get Work Item types in Project Area
```
ewm-cli get type --context _isel8HoTEeyoCYEXLNJaZA
```

### Get workflows in Project Area
```
ewm-cli get workflow --context _isel8HoTEeyoCYEXLNJaZA
```

### Get possible actions for a workflow in Project Area
```
ewm-cli get action --context _isel8HoTEeyoCYEXLNJaZA --workflow com.ibm.team.workitem.taskWorkflow
```

### Resolve task work item
```
ewm-cli resolve task --action com.ibm.team.workitem.taskWorkflow.action.complete --taskID 112
```

### Generate an HTML report of all defects created during checks
```
ewm-cli create defect-report
```