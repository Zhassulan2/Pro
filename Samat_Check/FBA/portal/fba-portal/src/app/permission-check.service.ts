import { Injectable } from '@angular/core';
import { OAuthService } from "angular-oauth2-oidc";

const roleConfig = 
{
  "cashier":["main-page"],
  "stockman":["main-page","products-receipt","category","products"],
  "accountant":["main-page","products-receipt","products-relocation","products","points","company","category","product-action"],
  "systemroot":["*"],
  "owner":["*"]
};

@Injectable()
export class PermissionCheckService {

  constructor(private oauthService: OAuthService) { }

  checkPermissionOnRoute(element: string) : boolean {
    if ((roleConfig[this.userRole][0] == "*") || (roleConfig[this.userRole].indexOf(element,0) > 0)) {
      return true;
    } else {
      return false;
    }
  }

  get userRole() {
    var claims = this.oauthService.getIdentityClaims();
    if (!claims) return null;
    return claims['Scopes'];
    //return "accountant"
  }
}
