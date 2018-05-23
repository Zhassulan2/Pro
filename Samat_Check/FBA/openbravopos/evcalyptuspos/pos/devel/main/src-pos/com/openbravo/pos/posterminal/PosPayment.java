/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.openbravo.pos.posterminal;

/**
 *
 * @author ikaryukin
 */

import java.io.*;
import com.sun.jna.Library;
import com.sun.jna.Native;
import java.util.HashMap;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.commons.io.IOUtils;
import com.openbravo.pos.forms.AppProperties;


public class PosPayment {
    
    AppProperties appProps;
    
    public PosPayment(AppProperties props) {
        appProps = props;        
    }
    
    
    public String makePaymentArcus(String amount) { 
        String result = null;        
        String[] cmd = { appProps.getProperty("payment.arcus")+"\\CommandLineTool\\bin\\CommandLineTool.exe", "/o1", "/a"+String.valueOf(amount), "/c398"};
        
        try {
            Process p = Runtime.getRuntime().exec(cmd);
            p.waitFor();           
                    
            try(FileInputStream inputStream = new FileInputStream(appProps.getProperty("payment.arcus")+"\\rc.out")) {     
                result = IOUtils.toString(inputStream);
                result = result.trim();
            }
        }
        catch(Exception e) {
            System.out.println(e.toString());
        }
                
    return result;            
    }
    
    public String getPaymentCodeDescr(String code) {
        String codeDescription = "";
        //Map<String, String> resCodes = new HashMap<>();
        try {         
            try(BufferedReader br = new BufferedReader(new FileReader(new File(appProps.getProperty("payment.arcus")+"\\INI\\rc_res_unicode.ini")))) {
                for(String line; (line = br.readLine()) != null; ) {
                    //resCodes.put(line.split("=")[0], line.split("=")[1]);
                    if ( (line.split("=")[0]).equals(code) )
                        codeDescription = line.split("=")[1];
                }                
            }
        } catch (IOException ex) {
            Logger.getLogger(PosPayment.class.getName()).log(Level.SEVERE, null, ex);
        }
       return codeDescription;
    }
    
}
