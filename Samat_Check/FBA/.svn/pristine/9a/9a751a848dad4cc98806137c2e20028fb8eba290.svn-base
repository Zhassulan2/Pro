import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AlertModule, BsDatepickerModule } from 'ngx-bootstrap';

import { AppComponent } from './app.component';
import { PasswordFlowLoginComponent } from './password-flow-login/password-flow-login.component';
import { AppRoutingModule } from './/app-routing.module';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HttpClient } from '@angular/common/http';
import { AuthConfig, JwksValidationHandler, OAuthModule, ValidationHandler } from 'angular-oauth2-oidc';
import { PointsComponent } from './points/points.component';
import { ProductsComponent } from './products/products.component';
import { MainPageComponent } from './main-page/main-page.component';
import { ChartModule } from 'angular-highcharts';
import { RegistrationComponent } from './registration/registration.component';
import { UserInfoComponent } from './user-info/user-info.component';
import { ProductActionComponent } from './product-action/product-action.component';
import { PermissionCheckService } from './permission-check.service';
import { RestService } from './rest.service';
import { StaffComponent } from './staff/staff.component';
import { OcticonDirective } from './custom-directives/octicon.directive';
import { TranslateModule, TranslateLoader } from "@ngx-translate/core";
import { TranslateHttpLoader } from "@ngx-translate/http-loader";
import { CategoryComponent } from './category/category.component';
import { CompanyComponent } from './company/company.component';
import { DeviceComponent } from './device/device.component';
import { SupplierComponent } from './supplier/supplier.component';
import { ProductsReceiptComponent } from './products-receipt/products-receipt.component';
import { ProductsRelocationComponent } from './products-relocation/products-relocation.component';
import { StaffPointComponent } from './staff-point/staff-point.component';
import { DiscountComponent } from './discount/discount.component';
import { ImportDataComponent } from './import-data/import-data.component';
import { ProductRefuseComponent } from './product-refuse/product-refuse.component';
import { InventoryComponent } from './inventory/inventory.component';
import { AlertComponent } from './alert/alert.component';
import { Alerts } from './alerts.model';
export function HttpLoaderFactory(httpClient: HttpClient) {
  return new TranslateHttpLoader(httpClient, "./assets/i18n/", ".json");
}

@NgModule({
  declarations: [
    AppComponent,
    PasswordFlowLoginComponent,
    PointsComponent,
    ProductsComponent,
    MainPageComponent,
    RegistrationComponent,
    UserInfoComponent,
    StaffComponent,
    OcticonDirective,
    CategoryComponent,
    CompanyComponent,
    DeviceComponent,
    SupplierComponent,
    ProductsReceiptComponent,
    ProductActionComponent,
    ProductsRelocationComponent,
    StaffPointComponent,
    DiscountComponent,
    ImportDataComponent,
    ProductRefuseComponent,
    InventoryComponent,
    AlertComponent
  ],
  imports: [
    ChartModule,
    AppRoutingModule,
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    OAuthModule.forRoot(),
    AlertModule.forRoot(),
    BsDatepickerModule.forRoot(),
    TranslateModule.forRoot({
      loader: {
        provide: TranslateLoader,
        useFactory: HttpLoaderFactory,
        deps: [HttpClient]
      }
    })
  ],
  providers: [RestService, PermissionCheckService, Alerts],
  bootstrap: [AppComponent]
})
export class AppModule { }
