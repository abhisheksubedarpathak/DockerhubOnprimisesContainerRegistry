# Copyright 2016-2017 VMware, Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#	http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License

*** Settings ***
Documentation  Longevity
Resource  ../../resources/Util.robot

*** Variables ***
${HARBOR_URL}  https://${ip}
${SSH_USER}  root
${SSH_PWD}  root1234
${HARBOR_PASSWORD}  Harbor12345

*** Keywords ***
Longevity setup
    Run Keyword  CA setup
    Run Keyword  Prepare Docker Cert  ${ip}
    Run Keyword  Start Docker Daemon Locally

CA setup
    Open Connection    ${ip}
    Login    ${SSH_USER}    ${SSH_PWD}
    SSHLibrary.Get File  /data/ca_download/ca.crt
    Close All Connections
    Run  mv ca.crt harbor_ca.crt
    Generate Certificate Authority For Chrome  ${HARBOR_PASSWORD}	

Regression Test With DB
    [Arguments]  ${HARBOR_URL}
    Run Keyword And Continue On Failure  Exe Regression Test Cases  ${HARBOR_URL}

Exe Regression Test Cases
    [Arguments]  ${HARBOR_URL}
    
    # New user, new project, push image, pull image
    Init Chrome Driver
    ${d}=    Get Current Date    result_format=%m%s
    Create An New Project With New User  url=${HARBOR_URL}  username=tester${d}  email=tester${d}@vmware.com  realname=tester${d}  newPassword=Test1@34  comment=harbor  projectname=project${d}  public=false
    Push image  ${ip}  tester${d}  Test1@34  project${d}  busybox:latest
    Go Into Project  project${d}
    Wait Until Page Contains  project${d}/busybox
    Pull image  ${ip}  tester${d}  Test1@34  project${d}  busybox:latest
    Close Browser

*** Test Cases ***
Longevity
    Run Keyword  Longevity setup
    # Each loop should take between 1 and 2 hours
    :FOR  ${idx}  IN RANGE  0  48
    \   ${rand}=  Evaluate  random.randint(10, 50)  modules=random
    \   Log To Console  \nLoop: ${idx}
    \   Repeat Keyword  ${rand} times  Regression Test With DB  ${HARBOR_URL}