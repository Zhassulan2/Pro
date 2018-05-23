import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { ActionType } from "../models/actionType";
@Component({
  selector: 'app-action-type',
  templateUrl: './action-type.component.html',
  styleUrls: ['./action-type.component.css']
})
export class ActionTypeComponent implements OnInit {
  actionTypes: Array<ActionType> = [];
  selectedActionType: ActionType;
  clientSecret: string;
  actionTypesScope: boolean = false;
  searchScope: boolean = false;
  paginationActionTypes: boolean = false;
  actionTypePages = [];
  actionTypeSize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService
  ) { }

  ngOnInit() {
    this.getActionTypes();
  }

  refreshActionType() {
    this.getActionTypes();
  }

  onSelect(actionType: ActionType): void {
    this.selectedActionType = actionType;
  }

  getPagination() {
    this.rest.getActionTypesCount().subscribe(
      res => {
        let actionTypeCount = parseInt(res.toString());
        if (actionTypeCount > this.actionTypeSize) {
          this.actionTypePages = [];
          for (var i = 0; i < Math.ceil(actionTypeCount / this.actionTypeSize); i++) {
            this.actionTypePages.push(i + 1)
          }
          this.paginationActionTypes = true;
        } else {
          this.actionTypePages = [];
          this.paginationActionTypes = false;
        }
      }
    )
  }

  getActionTypes(page?: number) {
    this.actionTypes = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getActionTypes(page, 10).subscribe(
      (actionTypes) => {
        this.actionTypes=actionTypes;
        this.actionTypesScope = true;
        this.getPagination();
      }
    );
  }

  createActionType(name: string) {
    let postActionType: ActionType = {Name: name};
    this.rest.postActionTypes(postActionType).subscribe(
      (postActionType) => {
        this.getActionTypes();
      }
    );
  }

  updateActionType() {
    this.rest.putActionTypes(this.selectedActionType).subscribe(
      (selectedActionType) => {
        this.getActionTypes();
      }
    );
  }

  deleteActionType() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedActionType.Name + "?")) {
      this.rest.deleteActionTypes(this.selectedActionType).subscribe(
        (selectedActionType) => {
          this.getActionTypes();
        }
      );
    }
  }

}
