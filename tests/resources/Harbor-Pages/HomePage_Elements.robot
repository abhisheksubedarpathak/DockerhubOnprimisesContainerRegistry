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
Documentation  This resource provides any keywords related to the Harbor private registry appliance

*** Variables ***
${sign_up_for_an_account_xpath}  /html/body/harbor-app/harbor-shell/clr-main-container/div/div/sign-in/div/form/div[1]/a
${username_xpath}  //*[@id="username"]
${email_xpath}  //*[@id="email"]
${realname_xpath}  //*[@id="realname"]
${newPassword_xpath}  //*[@id="newPassword"]
${confirmPassword_xpath}  //*[@id="confirmPassword"]
${comment_xpath}  //*[@id="comment"]
${signup_xpath}  /html/body/harbor-app/harbor-shell/clr-main-container/div/div/sign-in/sign-up/clr-modal/div/div[1]/div/div[1]/div/div[3]/button[2]

