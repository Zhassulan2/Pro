import { Component, OnInit } from '@angular/core';
import { Alerts } from '../alerts.model';

@Component({
  selector: 'app-alert',
  templateUrl: './alert.component.html',
  styleUrls: ['./alert.component.css']
})
export class AlertComponent implements OnInit {

  constructor(public alertScope: Alerts) { }

  ngOnInit() {
  }
}
