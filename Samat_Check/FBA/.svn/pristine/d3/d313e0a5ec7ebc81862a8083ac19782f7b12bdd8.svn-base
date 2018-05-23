//    Openbravo POS is a point of sales application designed for touch screens.
//    Copyright (C) 2008-2009 Openbravo, S.L.
//    http://www.openbravo.com/product/pos
//
//    This file is part of Openbravo POS.
//
//    Openbravo POS is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    Openbravo POS is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with Openbravo POS.  If not, see <http://www.gnu.org/licenses/>.

package com.openbravo.pos.payment;

import com.openbravo.pos.forms.AppProperties;
import com.openbravo.pos.posterminal.PosPayment;
import java.text.DecimalFormat;


public class PaymentGatewayExt implements PaymentGateway {  
    /** Creates a new instance of PaymentGatewayExt */
    
    AppProperties appProps;
    
    public PaymentGatewayExt(AppProperties props) {
        appProps = props;
    }
  
    public void execute(PaymentInfoMagcard payinfo) {
        
        String sReader = appProps.getProperty("payment.gateway");

        if ("ForteBank".equals(sReader)) {
            PosPayment pospay = new PosPayment(appProps);
            DecimalFormat df = new DecimalFormat("0.00"); 
            String amount = df.format(payinfo.getTotal()).replace(".", "").replace(",", "");

            //JOptionPane.showMessageDialog(null, "["+amount+"]");

            String payResult = pospay.makePaymentArcus(amount);

            //JOptionPane.showMessageDialog(null, "["+payResult+"]");

            if(Integer.parseInt(payResult) < 10)        
                payinfo.paymentOK("Успешно", payinfo.getTransactionID() , "");
            else {
                String errorText = "Код ошибки: "+payResult+System.lineSeparator()+"Описание ошибки: "+pospay.getPaymentCodeDescr(payResult);
                payinfo.paymentError("ОШИБКА ПРИ ОПЛАТЕ", errorText);
            }
        }
        else {
            payinfo.paymentOK("OK", payinfo.getTransactionID() , "");
        }
    }
}
