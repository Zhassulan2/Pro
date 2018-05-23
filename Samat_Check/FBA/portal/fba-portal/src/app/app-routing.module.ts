import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PasswordFlowLoginComponent } from './password-flow-login/password-flow-login.component';
import { PointsComponent } from './points/points.component';
import { ProductsComponent } from './products/products.component';
import { DiscountComponent } from './discount/discount.component';
import { ProductsReceiptComponent } from './products-receipt/products-receipt.component';
import { ProductsRelocationComponent } from './products-relocation/products-relocation.component';
import { MainPageComponent } from './main-page/main-page.component';
import { RegistrationComponent } from './registration/registration.component';
import { StaffComponent } from './staff/staff.component';
import { ProductRefuseComponent } from './product-refuse/product-refuse.component';
import { CompanyComponent } from './company/company.component';
import { StaffPointComponent } from './staff-point/staff-point.component';
import { CategoryComponent } from './category/category.component';
import { DeviceComponent } from './device/device.component';
import { SupplierComponent } from './supplier/supplier.component';
import { ProductActionComponent } from './product-action/product-action.component';
import { ImportDataComponent } from './import-data/import-data.component';
import { InventoryComponent } from './inventory/inventory.component';
import { RouteGuard } from './route.guard';

const routes: Routes = [
  { 
    path: 'main-page', 
    component: MainPageComponent,
    canActivate: [RouteGuard] 
  },
  { 
    path: 'password-flow-login', 
    component: PasswordFlowLoginComponent 
  },
  {
    path: 'import-data', 
    component: ImportDataComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'points', 
    component: PointsComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'staff', 
    component: StaffComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'device', 
    component: DeviceComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'products', 
    component: ProductsComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'inventory', 
    component: InventoryComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'discount', 
    component: DiscountComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'staff-point', 
    component: StaffPointComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'products-receipt', 
    component: ProductsReceiptComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'products-relocation', 
    component: ProductsRelocationComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'product-refuse', 
    component: ProductRefuseComponent,
    canActivate: [RouteGuard]
  },
  
  {
    path: 'company', 
    component: CompanyComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'category', 
    component: CategoryComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'supplier', 
    component: SupplierComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'product-action', 
    component: ProductActionComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'product-action/:id', 
    component: ProductActionComponent,
    canActivate: [RouteGuard]
  },
  {
    path: 'registration', 
    component: RegistrationComponent 
  },
  {
    path: '**', 
    component: PasswordFlowLoginComponent 
  }
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ],
  declarations: [],
  providers: [RouteGuard]
})
export class AppRoutingModule { }
