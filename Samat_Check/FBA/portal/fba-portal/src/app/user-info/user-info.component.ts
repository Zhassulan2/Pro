import { Component, OnInit } from '@angular/core';
import { OAuthService } from "angular-oauth2-oidc";
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.css']
})
export class UserInfoComponent implements OnInit {

  constructor(
    private oauthService: OAuthService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  get userId() {
    var claims = this.oauthService.getIdentityClaims();
    if (!claims) return null;
    return claims['ID'];
  }

  get givenName() {
    var claims = this.oauthService.getIdentityClaims();
    if (!claims) return null;
    return claims['Name'];
  }

  logout() {
    this.oauthService.logOut(true);
    this.router.navigate(["password-flow-login"]);
  }
}
