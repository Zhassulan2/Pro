import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { Device } from "../models/device";
import { Point } from "../models/point";
import { PermissionCheckService } from '../permission-check.service';

@Component({
  selector: 'app-device',
  templateUrl: './device.component.html',
  styleUrls: ['./device.component.css']
})
export class DeviceComponent implements OnInit {
  //New Device
  newName: string;
  newPointID: string = "";
  
  testname: string;
  devices: Array<Device> = [];
  points: Array<Point> = [];
  selectedDevice: Device;
  clientSecret: string;
  devicesScope: boolean = false;
  searchScope: boolean = false;
  paginationDevices: boolean = false;
  devicePages = [];
  deviceSize: number = 10;
  currentPage: number = 1;
  
  constructor(
    private rest: RestService,
    private check: PermissionCheckService
  ) { }

  ngOnInit() {
    this.getDevices();
    this.getPoints();
  }

  refreshDevice() {
    this.getDevices();
  }

  onSelect(device: Device): void {
    this.selectedDevice = device;
  }

  getPoints() {
    this.points = [];
    this.rest.getPointsCount().subscribe(
      res => {
        let pointsCount = parseInt(res.toString());
        this.rest.getPoints(1, pointsCount).subscribe(
          (points) => {
            this.points=points;
          }
        );
      }
    )
  }

  getPagination() {
    this.rest.getDevicesCount().subscribe(
      res => {
        let deviceCount = parseInt(res.toString());
        if (deviceCount > this.deviceSize) {
          this.devicePages = [];
          for (var i = 0; i < Math.ceil(deviceCount / this.deviceSize); i++) {
            this.devicePages.push(i + 1)
          }
          this.paginationDevices = true;
        } else {
          this.devicePages = [];
          this.paginationDevices = false;
        }
      }
    )
  }

  getDevices(page?: number) {
    this.devices = [];
    if (page) {
      this.currentPage = page;
    }
    this.rest.getDevices(page, 10).subscribe(
      (devices) => {
        this.devices=devices;
        this.devicesScope = true;
        this.getPagination();
      }
    );
  }

  createDevice() {
    let postDevice: Device = {Name: this.newName, PointID: this.newPointID};
    this.rest.postDevices(postDevice).subscribe(
      (postDevice) => {
        this.getDevices();
        this.newName = null;
      }
    );
  }

  updateDevice() {
    this.rest.putDevices(this.selectedDevice).subscribe(
      (selectedDevice) => {
        this.getDevices();
      }
    );
  }

  deleteDevice() {
    if (confirm("Вы уверены что хотите удалить точку " + this.selectedDevice.Name + "?")) {
      this.rest.deleteDevices(this.selectedDevice).subscribe(
        (selectedDevice) => {
          this.getDevices();
        }
      );
    }
  }
}

