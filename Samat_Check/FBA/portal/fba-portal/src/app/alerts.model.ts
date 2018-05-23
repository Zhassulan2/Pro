import { Injectable } from "@angular/core";
import { Alert } from "./models/alert";

@Injectable()
export class Alerts {
    public alerts: Alert[] = [];
    
    pushAlert(typeMessage: string, message: string) {
        let alertMessage : Alert = {
            type: typeMessage,
            msg: message
        }
        //if(!this.alerts.find(x=>x.msg == message)) {
            this.alerts.push(alertMessage);
        //}
    }
}