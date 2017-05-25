*** Settings ***
Documentation  Harbor BATs
Resource  ../../resources/Util.robot
Suite Setup  Install Harbor to Test Server
Default Tags  BAT

*** Test Cases ***
Test Case - Create An New Project
    ${d}=    Get Current Date    result_format=%m%s
    Start Selenium Standalone Server Locally
    Create An New User  username=tester${d}  email=tester${d}@vmware.com  realname=harbortest  newPassword=Test1@34  comment=harbortest
    Sign In Harbor  tester${d}  Test1@34
    Create An New Project  test${d}
    Close Browser

Test Case - Create An New User
    ${d}=    Get Current Date    result_format=%m%s
    Create An New User  username=tester${d}  email=tester${d}@vmware.com  realname=harbortest  newPassword=Test1@34  comment=harbortest

Test Case - Sign With Admin
    Sign In Harbor  %{HARBOR_ADMIN}  %{HARBOR_PASSWORD}
