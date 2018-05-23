import { Component, OnInit } from '@angular/core';
import { OAuthService } from 'angular-oauth2-oidc';
import { authPasswordFlowConfig } from '../auth-password-flow.config';
import { Router } from '@angular/router';
import { User } from "../models/user";
import { Alerts } from "../alerts.model";
import { RestService } from '../rest.service';


@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {
  dismissible = true;

  userPhone: string;
  userName: string;
  validationCode: string;
  userPassword: string;
  //showPassSet: boolean = false;
  showPassSet: boolean = true;
  showValidationPhone: boolean = false;

  constructor(
    private oauthService: OAuthService,
    private router: Router,
    private rest: RestService,
    private alerts: Alerts
  ) { 
    this.oauthService.configure(authPasswordFlowConfig)
    this.oauthService.loadDiscoveryDocument();}

  ngOnInit() {
  }

  loginWithPassword(login: string, password: string) {
    this
        .oauthService
        .fetchTokenUsingPasswordFlowAndLoadUserProfile(login, password)
        .then(() => {
          this.router.navigate(["points"]);
        })
        .catch((err) => {
          this.alerts.pushAlert("danger", "ALERTS.LOGINERR");
        });
  }

  registrationUser() {
    if ((this.userPassword.length > 7) && (this.userPassword.length < 21)){
      let user: User = {Login: this.userPhone, Password: this.userPassword, Name: this.userName, Scopes: "owner"}
      this.rest.postUser(user,"read,write")
        .subscribe(
        (user) => {
          if (user) {
            this.loginWithPassword(this.userPhone, this.userPassword);
          } else {
            this.alerts.pushAlert("danger", "ALERTS.USERERR");
          }
        }
      );
    } else {
      this.alerts.pushAlert("danger", "ALERTS.PASSWORDERR");
    }
  }

  sendVerificationCode() {
    this.alerts.pushAlert("success", "ALERTS.PHONECODE");
    this.showValidationPhone = true;
  }

  checkVerificationCode() {
    this.showPassSet = true;
    this.showValidationPhone = false;
  }
}
