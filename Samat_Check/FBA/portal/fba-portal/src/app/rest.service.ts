import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders, HttpParams, HttpErrorResponse } from '@angular/common/http';
import { OAuthService } from "angular-oauth2-oidc";
import { Router } from '@angular/router';
import { Observable } from "rxjs/Observable";
import { of } from 'rxjs/observable/of';
import { catchError, map, tap } from 'rxjs/operators';

//Модели
import { Point } from "./models/point";
import { Device } from "./models/device";
import { StaffPoint } from "./models/staffPoint";
import { Discount } from "./models/discount";
import { DiscountType } from "./models/discountType";
import { City } from "./models/city";
import { ActionType } from "./models/actionType";
import { PaymentType } from "./models/paymentType";
import { Client } from "./models/client";
import { User } from "./models/user";
import { Tax } from "./models/tax";
import { Category } from "./models/category";
import { Staff } from "./models/staff";
import { Product } from "./models/product";
import { Company } from "./models/company";
import { Supplier } from "./models/supplier";
import { Role } from "./models/role";
import { ProductAction } from "./models/productAction";
import { ProductActionDetail } from "./models/productActionDetail";
import { Inventory } from "./models/inventory";
import { InventoryDetail } from "./models/inventoryDetail";
import { ProductStock } from "./models/productStock";
import { ImportProduct } from "./models/importProduct";
import { Remainder, IncomeSales } from "./models/reports";
import { Alerts } from "./alerts.model";


const portalUrl = 'http://localhost:8080';
const oauthUrl = 'http://localhost:9096';
/*
const portalUrl = 'http://192.168.8.105:8080';
const oauthUrl = 'http://192.168.8.105:9096';
*/
@Injectable()
export class RestService {
    constructor(
        private http: HttpClient,
        private oauthService: OAuthService,
        private router: Router,
        private alerts: Alerts
    ) {}

    //Oauth2
    postUser(user: User, roles: string) {
        return this.http.post<User>(oauthUrl + "/registrationuser?roles="+roles, user)
        .pipe(
            catchError(this.handleError<User>())
        );
    }

    getClientInfo(client: Client): Observable<Client> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get<Client>(portalUrl + "/point/" + client.ID + "/authorize", { headers })
        .pipe(
            catchError(this.handleError<Client>())
        );
    }

    //Role
    getRoles(page?: number, size?: number): Observable<Role[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/role";
        if  (page && size) {
            getUrl = portalUrl + "/role?page=" + page + "&size=" + size;
        }
        return this.http.get<Role[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    //Points
    getPointsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/points/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getPoints(page?: number, size?: number, company?: string): Observable<Point[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/points";
        if ((page && size) || company) {
            getUrl = getUrl + "?"; 
            if  (page && size) {
                getUrl = getUrl + "page=" + page + "&size=" + size;
            }
            if  (company) {
                getUrl = getUrl + "&companyid=" + company;
            }
        }
        return this.http.get<Point[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postPoints(point: Point): Observable<Point> {
        point.City = null;
        point.Company = null;
        point.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Point>(portalUrl+"/point", point, { headers })
        .pipe(
            catchError(this.handleError<Point>())
        );
    }

    putPoints(point: Point): Observable<any> {
        point.Company = null;
        point.City = null;
        point.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/point", point, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deletePoints(point: Point): Observable<Point> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Point>(portalUrl + "/point/" + point.ID, { headers })
        .pipe(
            catchError(this.handleError<Point>())
        );
    }

    //City
    getCitiesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/cities/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getCities(page?: number, size?: number): Observable<City[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/city";
        if  (page && size) {
            getUrl = portalUrl + "/city?page=" + page + "&size=" + size;
        }
        return this.http.get<City[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postCities(city: City): Observable<City> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<City>(portalUrl+"/city", city, { headers })
        .pipe(
            catchError(this.handleError<City>())
        );
    }

    putCities(city: City): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/city", city, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteCities(city: City): Observable<City> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<City>(portalUrl + "/city/" + city.ID, { headers })
        .pipe(
            catchError(this.handleError<City>())
        );
    }

    //Discount
    getDiscountsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/discounts/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getDiscounts(page?: number, size?: number): Observable<Discount[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/discount";
        if  (page && size) {
            getUrl = portalUrl + "/discount?page=" + page + "&size=" + size;
        }
        return this.http.get<Discount[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postDiscounts(discount: Discount): Observable<Discount> {
        discount.DiscountType = null;
        discount.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Discount>(portalUrl+"/discount", discount, { headers })
        .pipe(
            catchError(this.handleError<Discount>())
        );
    }

    putDiscounts(discount: Discount): Observable<any> {
        discount.DiscountType = null;
        discount.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/discount", discount, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteDiscounts(discount: Discount): Observable<Discount> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Discount>(portalUrl + "/discount/" + discount.ID, { headers })
        .pipe(
            catchError(this.handleError<Discount>())
        );
    }

    //PaymentType
    getPaymentTypesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/paymenttypes/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getPaymentTypes(page?: number, size?: number): Observable<PaymentType[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/paymenttype";
        if  (page && size) {
            getUrl = portalUrl + "/paymenttype?page=" + page + "&size=" + size;
        }
        return this.http.get<PaymentType[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    getPaymentTypeByShortName(shortName: string): Observable<PaymentType> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/paymenttype?shortname="+shortName;
        
        return this.http.get<PaymentType>(getUrl, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    postPaymentTypes(paymentType: PaymentType): Observable<PaymentType> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<PaymentType>(portalUrl+"/paymenttype", paymentType, { headers })
        .pipe(
            catchError(this.handleError<PaymentType>())
        );
    }

    putPaymentTypes(paymentType: PaymentType): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/paymenttype", paymentType, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deletePaymentTypes(paymentType: PaymentType): Observable<PaymentType> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<PaymentType>(portalUrl + "/paymenttype/" + paymentType.ID, { headers })
        .pipe(
            catchError(this.handleError<PaymentType>())
        );
    }

    //DiscountType
    getDiscountTypesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/discounttypes/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getDiscountTypes(page?: number, size?: number): Observable<DiscountType[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/discounttype";
        if  (page && size) {
            getUrl = portalUrl + "/discounttype?page=" + page + "&size=" + size;
        }
        return this.http.get<DiscountType[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    //ActionType
    getActionTypesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/actiontypes/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getActionTypes(page?: number, size?: number): Observable<ActionType[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/actiontype";
        if  (page && size) {
            getUrl = portalUrl + "/actiontype?page=" + page + "&size=" + size;
        }
        return this.http.get<ActionType[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    getActionTypeByShortName(shortName: string): Observable<ActionType> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/actiontype?shortname="+shortName;
        return this.http.get<ActionType>(getUrl, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    postActionTypes(actionType: ActionType): Observable<ActionType> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<ActionType>(portalUrl+"/actiontype", actionType, { headers })
        .pipe(
            catchError(this.handleError<ActionType>())
        );
    }

    putActionTypes(actionType: ActionType): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/actiontype", actionType, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteActionTypes(actionType: ActionType): Observable<ActionType> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<ActionType>(portalUrl + "/actiontype/" + actionType.ID, { headers })
        .pipe(
            catchError(this.handleError<ActionType>())
        );
    }

    //Tax
    getTaxesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/taxs/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getTaxByShortName(shortName: string): Observable<Tax> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/tax?shortname="+shortName;
        
        return this.http.get<Tax>(getUrl, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }
    getTaxes(page?: number, size?: number): Observable<Tax[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/tax";
        if  (page && size) {
            getUrl = portalUrl + "/tax?page=" + page + "&size=" + size;
        }
        return this.http.get<Tax[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postTaxes(tax: Tax): Observable<Tax> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Tax>(portalUrl+"/tax", tax, { headers })
        .pipe(
            catchError(this.handleError<Tax>())
        );
    }

    putTaxes(tax: Tax): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/tax", tax, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteTaxes(tax: Tax): Observable<Tax> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Tax>(portalUrl + "/tax/" + tax.ID, { headers })
        .pipe(
            catchError(this.handleError<Tax>())
        );
    }

    //Category
    getCategoriesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/categories/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getCategories(page?: number, size?: number): Observable<Category[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/category";
        if  (page && size) {
            getUrl = portalUrl + "/category?page=" + page + "&size=" + size;
        }
        return this.http.get<Category[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postCategories(category: Category): Observable<Category> {
        category.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Category>(portalUrl+"/category", category, { headers })
        .pipe(
            catchError(this.handleError<Category>())
        );
    }

    putCategories(category: Category): Observable<any> {
        category.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/category", category, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteCategories(category: Category): Observable<Category> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Category>(portalUrl + "/category/" + category.ID, { headers })
        .pipe(
            catchError(this.handleError<Category>())
        );
    }

    //Staff
    getStaffsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/staffs/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getStaffs(page?: number, size?: number): Observable<Staff[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/staff";
        if  (page && size) {
            getUrl = portalUrl + "/staff?page=" + page + "&size=" + size;
        }
        return this.http.get<Staff[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postStaffs(staff: Staff): Observable<Staff> {
        staff.Role = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Staff>(portalUrl+"/staff", staff, { headers })
        .pipe(
            catchError(this.handleError<Staff>())
        );
    }

    putStaffs(staff: Staff): Observable<any> {
        staff.Role = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/staff", staff, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteStaffs(staff: Staff): Observable<Staff> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Staff>(portalUrl + "/staff/" + staff.ID, { headers })
        .pipe(
            catchError(this.handleError<Staff>())
        );
    }

    //Products
    getProductAvgPrice(productID: string, pointID: string): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/product/"+ productID +"/averageprice/" + pointID, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getProductLastPrice(productID: string, pointID: string): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/product/"+ productID +"/lastprice/" + pointID, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getProductsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/products/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }


    getProducts(page?: number, size?: number, name?: string, barcode?: string): Observable<Product[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/product";
        if ((page && size) || name || barcode) {
            getUrl = getUrl + "?"; 
            if  (page && size) {
                getUrl = getUrl + "page=" + page + "&size=" + size;
            }
            if  (name) {
                getUrl = getUrl + "&name=" + name;
            }
            if  (barcode) {
                getUrl = getUrl + "&barcode=" + barcode;
            }
        }
        return this.http.get<Product[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postProducts(product: Product): Observable<Product> {
        product.Category = null;
        product.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Product>(portalUrl+"/product", product, { headers })
        .pipe(
            catchError(this.handleError<Product>())
        );
    }

    putProducts(product: Product): Observable<any> {
        product.Category = null;
        product.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/product", product, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteProducts(product: Product): Observable<Product> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Product>(portalUrl + "/product/" + product.ID, { headers })
        .pipe(
            catchError(this.handleError<Product>())
        );
    }

    
    //Company
    getCompaniesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/companies/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getCompanies(page?: number, size?: number): Observable<Company[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/company";
        if  (page && size) {
            getUrl = portalUrl + "/company?page=" + page + "&size=" + size;
        }
        return this.http.get<Company[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postCompanies(company: Company): Observable<Company> {
        company.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Company>(portalUrl+"/company", company, { headers })
        .pipe(
            catchError(this.handleError<Company>())
        );
    }

    putCompanies(company: Company): Observable<any> {
        company.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/company", company, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteCompanies(company: Company): Observable<Company> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Company>(portalUrl + "/company/" + company.ID, { headers })
        .pipe(
            catchError(this.handleError<Company>())
        );
    }
    
    //Supplier
    getSuppliersCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/suppliers/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getSuppliers(page?: number, size?: number): Observable<Supplier[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/supplier";
        if  (page && size) {
            getUrl = portalUrl + "/supplier?page=" + page + "&size=" + size;
        }
        return this.http.get<Supplier[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postSuppliers(supplier: Supplier): Observable<Supplier> {
        supplier.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Supplier>(portalUrl+"/supplier", supplier, { headers })
        .pipe(
            catchError(this.handleError<Supplier>())
        );
    }

    putSuppliers(supplier: Supplier): Observable<any> {
        supplier.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/supplier", supplier, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteSuppliers(supplier: Supplier): Observable<Supplier> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Supplier>(portalUrl + "/supplier/" + supplier.ID, { headers })
        .pipe(
            catchError(this.handleError<Supplier>())
        );
    }

    //ProductAction
    getProductActionsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/productactions/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getProductActions(page?: number, size?: number): Observable<ProductAction[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/productaction";
        if  (page && size) {
            getUrl = portalUrl + "/productaction?page=" + page + "&size=" + size;
        }
        return this.http.get<ProductAction[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postProductAction(productAction: ProductAction): Observable<ProductAction> {
        productAction.Point = null;
        productAction.Staff = null;
        productAction.WorkSession = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<ProductAction>(portalUrl+"/productaction", productAction, { headers })
        .pipe(
            catchError(this.handleError<ProductAction>())
        );
    }

    //ProductActionDetail
    getProductActionDetailsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/productactiondetails/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getProductActionDetails(page?: number, size?: number, productActionID?: string): Observable<ProductActionDetail[]> {
        let getUrl = portalUrl + "/productactiondetail";
        if ((page && size) || productActionID) {
            getUrl = getUrl + "?"; 
            if  (page && size) {
                getUrl = getUrl + "page=" + page + "&size=" + size;
            }
            if  (productActionID) {
                getUrl = getUrl + "&productActionID=" + productActionID;
            }
        }
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get<ProductActionDetail[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postProductActionDetails(productActionDetails: ProductActionDetail[]): Observable<ProductActionDetail[]> {
        let postProductActionDetails: ProductActionDetail[] = productActionDetails;
        for (var i = 0; i < productActionDetails.length; i++) {
            postProductActionDetails[i].ActionType = null;
            postProductActionDetails[i].PaymentType = null;
            postProductActionDetails[i].ProductAction = null;
            postProductActionDetails[i].Product = null;
            postProductActionDetails[i].Supplier = null;
            postProductActionDetails[i].Tax = null;
        }
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<ProductActionDetail[]>(portalUrl+"/productactiondetails", postProductActionDetails, { headers })
        .pipe(
            catchError(this.handleError<ProductActionDetail[]>())
        );
    }

    //ProductStock
    getProductStockCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/productstocks/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    postProductStock(productStock: ProductStock): Observable<ProductStock> {
        productStock.Point = null;
        productStock.Product = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<ProductStock>(portalUrl+"/productstock", productStock, { headers })
        .pipe(
            catchError(this.handleError<ProductStock>())
        );
    }

    getProductStock(page?: number, size?: number): Observable<ProductStock[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/productstock";
        if  (page && size) {
            getUrl = portalUrl + "/productstock?page=" + page + "&size=" + size;
        }
        return this.http.get<ProductStock[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    getProductStockByPID(pointID: string, page?: number, size?: number, name?: string, barcode?: string): Observable<ProductStock[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/productstock/" + pointID;
        if  (page && size) {
            getUrl = getUrl + "?page=" + page + "&size=" + size;
        }
        if  (name) {
            getUrl = getUrl + "&productname=" + name;
        }
        if  (barcode) {
            getUrl = getUrl + "&productbarcode=" + barcode;
        }
        return this.http.get<ProductStock[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    //StaffPoint
    getStaffPointsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/staffpoints/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getStaffPoints(page?: number, size?: number): Observable<StaffPoint[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/staffpoint";
        if  (page && size) {
            getUrl = portalUrl + "/staffpoint?page=" + page + "&size=" + size;
        }
        return this.http.get<StaffPoint[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postStaffPoint(staffPoint: StaffPoint): Observable<StaffPoint> {
        staffPoint.Point = null;
        staffPoint.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<StaffPoint>(portalUrl+"/staffpoint", staffPoint, { headers })
        .pipe(
            catchError(this.handleError<StaffPoint>())
        );
    }

    deleteStaffPoint(staffPoint: StaffPoint): Observable<StaffPoint> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<StaffPoint>(portalUrl + "/staffpoint/" + staffPoint.ID, { headers })
        .pipe(
            catchError(this.handleError<StaffPoint>())
        );
    }

    //Device
    getDevicesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/devices/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getDevices(page?: number, size?: number): Observable<Device[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/device";
        if  (page && size) {
            getUrl = portalUrl + "/device?page=" + page + "&size=" + size;
        }
        return this.http.get<Device[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postDevices(device: Device): Observable<Device> {
        device.Point = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Device>(portalUrl+"/device", device, { headers })
        .pipe(
            catchError(this.handleError<Device>())
        );
    }

    putDevices(device: Device): Observable<any> {
        device.Point = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/device", device, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteDevices(device: Device): Observable<Device> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Device>(portalUrl + "/device/" + device.ID, { headers })
        .pipe(
            catchError(this.handleError<Device>())
        );
    }

    //Inventory
    getInventoriesCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/inventories/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getInventories(page?: number, size?: number): Observable<Inventory[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/inventory";
        if  (page && size) {
            getUrl = portalUrl + "/inventory?page=" + page + "&size=" + size;
        }
        return this.http.get<Inventory[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postInventories(inventory: Inventory): Observable<Inventory> {
        inventory.Company = null;
        inventory.Point = null;
        inventory.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<Inventory>(portalUrl+"/inventory", inventory, { headers })
        .pipe(
            catchError(this.handleError<Inventory>())
        );
    }

    putInventories(inventory: Inventory): Observable<any> {
        inventory.Company = null;
        inventory.Point = null;
        inventory.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/inventory", inventory, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteInventories(inventory: Inventory): Observable<Inventory> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<Inventory>(portalUrl + "/inventory/" + inventory.ID, { headers })
        .pipe(
            catchError(this.handleError<Inventory>())
        );
    }

    //InventoryDetail
    getInventoryDetailsCount(): Observable<any> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.get(portalUrl+"/inventorydetails/count", { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    getInventoryDetails(inventory: string, page?: number, size?: number): Observable<InventoryDetail[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/inventorydetail?inventoryID=" + inventory;        
        if  (page && size) {
            getUrl = getUrl + "&page=" + page + "&size=" + size;
        }
        return this.http.get<InventoryDetail[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    postInventoryDetails(inventoryDetail: InventoryDetail): Observable<InventoryDetail> {
        inventoryDetail.Inventory = null;
        inventoryDetail.Product = null;
        inventoryDetail.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<InventoryDetail>(portalUrl+"/inventorydetail", inventoryDetail, { headers })
        .pipe(
            catchError(this.handleError<InventoryDetail>())
        );
    }

    putInventoryDetails(inventoryDetail: InventoryDetail): Observable<any> {
        inventoryDetail.Inventory = null;
        inventoryDetail.Product = null;
        inventoryDetail.Staff = null;
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.put(portalUrl + "/inventoryDetail", inventoryDetail, { headers })
        .pipe(
            catchError(this.handleError<any>())
        );
    }

    deleteInventoryDetails(inventoryDetail: InventoryDetail): Observable<InventoryDetail> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.delete<InventoryDetail>(portalUrl + "/inventoryDetail/" + inventoryDetail.ID, { headers })
        .pipe(
            catchError(this.handleError<InventoryDetail>())
        );
    }

    //ImportProduct
    postImportProduct(importProduct: ImportProduct[], point: string): Observable<ProductAction> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        return this.http.post<ProductAction>(portalUrl+"/import/product/" + point, importProduct, { headers })
        .pipe(
            catchError(this.handleError<ProductAction>())
        );
    }

    //Reports
    reportRemainder(): Observable<Remainder[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/report/remainder";
        return this.http.get<Remainder[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    reportIncomeSales(companyID: string, startDate: string, endDate: string): Observable<IncomeSales[]> {
        let headers = new HttpHeaders()
            .set("Authorization", "Bearer " + this.oauthService.getAccessToken());
        let getUrl = portalUrl + "/report/incomesales/" + companyID + "?startdate=" + startDate + "&enddate=" + endDate;
        return this.http.get<IncomeSales[]>(getUrl, { headers })
        .pipe(
            catchError(this.handleError([]))
        );
    }

    //Error
    private handleError<T> (result?: T) {
        return (error: HttpErrorResponse): Observable<T> => {
            if (error.status == 401) {
                this.alerts.pushAlert("warning", "ALERTS.AUTHORIZATION");        
                this.oauthService.logOut(true);
                this.router.navigate(["password-flow-login"]);         
            }
            if (error.status == 502) {
                this.alerts.pushAlert("danger", error.error);            
            }
            console.error(error);
            return of(result as T);
        };
    }
}