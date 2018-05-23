import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';
import { Remainder, IncomeSales } from "../models/reports";
import { Company } from "../models/company";
import { Point } from "../models/point";
import { Chart } from 'angular-highcharts';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.css']
})
export class MainPageComponent implements OnInit{
  remainderResult: Remainder[];
  remainderResultNull: boolean = false;
  incomeSalesResult: IncomeSales[];
  incomeSalesResultGraph = [];
  incomeSalesResultNull: boolean = false;
  repISStartDate: Date = new Date();
  repISEndDate: Date = new Date();
  repISCompany: string = "";
  //repISPoint: string = "";
  companies: Array<Company> = [];
  points: Array<Point> = [];
  chartIncomeSales: Chart;

  constructor(private rest: RestService) { }
 
  ngOnInit() {
      this.getCompanies();
      this.getRemainderReport();
  }

  getCompanies() {
    this.companies = [];
    this.rest.getCompaniesCount().subscribe(
      res => {
        let companiesCount = parseInt(res.toString());
        this.rest.getCompanies(1, companiesCount).subscribe(
          (companies) => {
            this.companies=companies;
          }
        );
      }
    )
  }

  /*
  getPoints() {
    this.points = [];
    this.rest.getPointsCount().subscribe(
      res => {
        let pointsCount = parseInt(res.toString());
        this.rest.getPoints(1, pointsCount, this.repISCompany).subscribe(
          (points) => {
            this.points=points;
          }
        );
      }
    )
  }
  */

  getRemainderReport() {
    this.remainderResultNull = false;
    this.remainderResult = null;
    this.rest.reportRemainder().subscribe(
        (remainderResult) => {
           if (remainderResult) {
                this.remainderResult=remainderResult;
            } else {
                this.remainderResultNull = true;
            }
        }
      );
  }

  getIncomeSalesReport() {
    this.incomeSalesResultNull = false;
    this.incomeSalesResultGraph = [];
    this.incomeSalesResult = null;
    let sDate = this.repISStartDate.toISOString().substr(0,10);
    let eDate = this.repISEndDate.toISOString().substr(0,10);
    this.rest.reportIncomeSales(this.repISCompany, sDate, eDate).subscribe(
        (incomeSalesResult) => {
            if (incomeSalesResult) {
                this.incomeSalesResult=incomeSalesResult;
                for (let i of this.incomeSalesResult) {
                    this.incomeSalesResultGraph.push([i.PointName, i.Costs]);
                }     
                this.generatePieChartIS(this.incomeSalesResultGraph);
            } else {
                this.incomeSalesResultNull = true;
            }
        }
      );
  }

  generatePieChartIS(seriesData) {
    this.chartIncomeSales = null; 
    this.chartIncomeSales = new Chart({  
        chart: {
          plotBackgroundColor: null,
          plotBorderWidth: null,
          plotShadow: false,
          type: 'pie'
        },
        title: {
            text: ''
        },
        tooltip: {
            pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
        },
        plotOptions: {
            pie: {
                allowPointSelect: true,
                cursor: 'pointer',
                dataLabels: {
                    enabled: true,
                    format: '<b>{point.name}</b>: {point.percentage:.1f} %'
                }
            }
        },
        series: [{
            data: seriesData
        }]
      });
  }

  /*
    import { Chart } from 'angular-highcharts';
    resultGetBalanceReports: Array<ResultGetBalanceReport> = [];
    BarReportCat = [];
    chart2: Chart;
    for (let res of this.resultGetGraphReports) {
        this.GraphReportData.push([res.NAME, res.TOTAL]);
        this.BarReportCat.push([res.NAME]);
    }     
    his.generateBarChart(this.GraphReportData, this.BarReportCat);

    generateBarChart(seriesData, catData) {
        this.chart2 = new Chart({  
            chart: {
            type: 'bar'
            },
            title: {
                text: 'Объём продаж'
            },
            xAxis: {
                categories: catData,
                title: {
                    text: null
                }
            },
            yAxis: {
                min: 0,
                title: {
                    text: 'Сумма продаж (KZT)',
                    align: 'high'
                },
                labels: {
                    overflow: 'justify'
                }
            },
            tooltip: {
                valueSuffix: ' KZT'
            },
            plotOptions: {
                bar: {
                    dataLabels: {
                        enabled: true
                    }
                }
            },
            credits: {
                enabled: false
            },
            series: [{
                name: 'Показатели сотрудников',
                data: seriesData
            }]
        });
    }
  */
}
