import { Component, OnInit } from '@angular/core';
import { OAuthService } from 'angular-oauth2-oidc';
import { authPasswordFlowConfig } from '../auth-password-flow.config';
import { ActivatedRoute } from '@angular/router';
import { Alert } from "../models/alert";

@Component({
  selector: 'app-password-flow-login',
  templateUrl: './password-flow-login.component.html',
  styleUrls: ['./password-flow-login.component.css']
})

export class PasswordFlowLoginComponent implements OnInit {
  userName: string;
  password: string;
  loginFailed: boolean = false;
  alerts: Array<Alert> = [];

  constructor(
    private oauthService: OAuthService,
    private route: ActivatedRoute) { 
    this.oauthService.configure(authPasswordFlowConfig)
    this.oauthService.loadDiscoveryDocument();
  }

  ngOnInit() {
  }

  get access_token() {
      return this.oauthService.getAccessToken();
  }

  get access_token_expiration() {
    return this.oauthService.getAccessTokenExpiration();
  }

  get givenName() {
      var claims = this.oauthService.getIdentityClaims();
      if (!claims) return null;
      return claims['Name'];
  }

  get userRole() {
    var claims = this.oauthService.getIdentityClaims();
    if (!claims) return null;
    return claims['Scopes'];
  }

  get userId() {
    var claims = this.oauthService.getIdentityClaims();
    if (!claims) return null;
    return claims['ID'];
  }

  loginWithPassword() {
    this
        .oauthService
        .fetchTokenUsingPasswordFlowAndLoadUserProfile(this.userName, this.password)
        .then(() => {
            console.debug('successfully logged in');
            this.loginFailed = false;
        })
        .catch((err) => {
            console.error('error logging in', err);
            this.loginFailed = true;
        });
  }

  logout() {
      this.oauthService.logOut(true);
  }
}
