*** Settings ***
Documentation  It's an demo case to deploy Harbor with Drone.
Resource  ../../resources/Util.robot
Suite Setup  Start Docker Daemon Locally
Default Tags  regression

*** Test Cases ***
Install Harbor to Test Server
    Log To Console  \nconfig harbor cfg
    Run Keywords  Config Harbor cfg
    Log To Console  \ncomplile and up harbor now
    Run Keywords  Compile and Up Harbor With Source Code
    ${rc}  ${output}=  Run And Return Rc And Output  docker ps
    Should Be Equal As Integers  ${rc}  0
    Log To Console  \n${output}

Test Harbor To Create An New User
    Start Selenium Standalone Server Locally
    ${d}=    Get Current Date    result_format=%m%s
    Create An New User  username=test${d}  email=test${d}@vmware.com  realname=harbortest  newPassword=Test1@34  comment=harbortest
