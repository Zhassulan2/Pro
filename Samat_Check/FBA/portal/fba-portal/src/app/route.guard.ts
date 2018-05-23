import { Injectable } from "@angular/core";
import { ActivatedRouteSnapshot, RouterStateSnapshot } from "@angular/router";
import { OAuthService } from "angular-oauth2-oidc";
import { Alerts } from "./alerts.model";
import { PermissionCheckService } from './permission-check.service';

@Injectable()
export class RouteGuard {
    private firstNavigation = true;
    constructor(
        private oauthService: OAuthService, 
        private check: PermissionCheckService,
        private alerts: Alerts
    ) { }
    canActivate(route: ActivatedRouteSnapshot,
        state: RouterStateSnapshot): boolean {
        if (this.oauthService.getAccessToken()) {
            if (this.check.checkPermissionOnRoute(route['url'][0].toString())) {
                return true;
            } else {
                this.alerts.pushAlert("danger","ALERTS.PERMISSIONROLE");
                return false;
            }
        } else {
            this.alerts.pushAlert("danger","ALERTS.PERMISSIONROUTE");
            return false;
        }
    }
}