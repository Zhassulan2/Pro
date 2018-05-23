import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { ProductAction } from "../models/productAction";
import { ProductActionDetail } from "../models/productActionDetail";
import { ActivatedRoute, Router } from '@angular/router';


@Component({
  selector: 'app-product-action',
  templateUrl: './product-action.component.html',
  styleUrls: ['./product-action.component.css']
})

export class ProductActionComponent implements OnInit {
  productActions: Array<ProductAction> = [];
  productActionDetails: Array<ProductActionDetail> = [];
  selectedProductAction: ProductAction;
  productActionsScope: boolean = false;
  productActionDetailsScope: boolean = false;
  
  productActionSize: number = 10;
  productActionDetailSize: number = 10;

  paginationProductActions: boolean = false;
  productActionPages = [];
  currentPagePA: number = 1;
  
  paginationProductActionDetails: boolean = false;
  productActionDetailPages = [];
  currentPagePAD: number = 1;
  padID = this.route.snapshot.paramMap.get('id');

  constructor(
    private rest: RestService,
    private route: ActivatedRoute,
    private router: Router
  ) { }

  ngOnInit() {
    if (this.padID) {
      this.getProductActionDetails();
    } else {
      this.getProductActions();
    }
  }

  onSelect(productAction: ProductAction): void {
    this.selectedProductAction = productAction;
  }

  getPaginationPA() {
    this.rest.getProductActionsCount().subscribe(
      res => {
        let productActionCount = parseInt(res.toString());
        if (productActionCount > this.productActionSize) {
          this.productActionPages = [];
          for (var i = 0; i < Math.ceil(productActionCount / this.productActionSize); i++) {
            this.productActionPages.push(i + 1)
          }
          this.paginationProductActions = true;
        } else {
          this.productActionPages = [];
          this.paginationProductActions = false;
        }
      }
    )
  }

  getProductActions(page?: number) {
    this.productActions = [];
    if (page) {
      this.currentPagePA = page;
    }
    this.rest.getProductActions(page, 10).subscribe(
      (productActions) => {
        this.productActions=productActions;
        this.productActionsScope = true;
        this.getPaginationPA();
      }
    );
  }

  getPaginationPAD() {
    this.rest.getProductActionDetailsCount().subscribe(
      res => {
        let productActionDetailCount = parseInt(res.toString());
        if (productActionDetailCount > this.productActionDetailSize) {
          this.productActionDetailPages = [];
          for (var i = 0; i < Math.ceil(productActionDetailCount / this.productActionDetailSize); i++) {
            this.productActionDetailPages.push(i + 1)
          }
          this.paginationProductActionDetails = true;
        } else {
          this.productActionDetailPages = [];
          this.paginationProductActionDetails = false;
        }
      }
    )
  }

  getProductActionDetails(page?: number) {
    this.productActionDetails = [];
    if (page) {
      this.currentPagePAD = page;
    }
    let productActionID: string;
    if (this.selectedProductAction) {
      productActionID = this.selectedProductAction.ID;
    } else {
      productActionID = this.padID;
    }
    this.rest.getProductActionDetails(page, 10, productActionID).subscribe(
      (productActionDetails) => {
        this.productActionDetails=productActionDetails;
        this.productActionDetailsScope = true;
        this.getPaginationPAD();
      }
    );
  }

  backToPA() {
    if (this.productActions.length == 0) {
      this.router.navigate(["product-action"]);
    }
    this.productActionDetails = [];
  }
}
