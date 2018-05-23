/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package net.ddns.evcalyptus.bonus;

import java.util.*;
import java.util.logging.Logger;
import com.openbravo.basic.BasicException;
import com.openbravo.data.loader.*;
import com.openbravo.pos.forms.BeanFactoryDataSingle;

/**
 *
 * @author user2
 */

/*
working notes and planning
=  read bonus config tables from BONUSCONFIGROLES and TEAMBONUSCONFIGHOSTS
=  create Maps with per role config settings
=  IndivBonus can be set for each role as % of individual sales

*/

public class DataLogicBonusConfig extends BeanFactoryDataSingle{
    Logger l = Logger.getLogger(getClass().getName());
	
    public DataLogicBonusConfig() {

	}

    protected Session s;
    
    public  static Map <String, Double> mapIndivBonus = new HashMap<>();
    public  static Map <String, Double> mapIndivPlan = new HashMap<>();
    //public  Map <String, Double> mapTeamBonusFromSales = new HashMap<>();
    public  static Map <String, Double> mapTeamBonusCoef = new HashMap<>();
    public static  double TB_FROM_SALES;
    public static  double TB_SESSION_PLAN;
    
    
    public void init (Session s){
        
        //mapIndivBonus = new HashMap<>();
        //mapTeamBonusFromSales = new HashMap<>();
        //mapTeamBonusCoef = new HashMap<>();
    	this.s = s;
    	int rows = 0;
    	try {
    		/*rows = (int) new PreparedSentence (s
    				, "SELECT COUNT(ID) FROM ROLES"
        			, null
        			, SerializerReadInteger.INSTANCE).find();
                */
                
                rows = countRolesRows();
                
                TB_FROM_SALES = (Double) new PreparedSentence (s
                                , "SELECT BONUS_FROM_SALES FROM ROLES WHERE ID = 'TEAMBONUSCONFIG'"
                                , null
                                , SerializerReadDouble.INSTANCE).find();
                
                TB_SESSION_PLAN = (Double) new PreparedSentence (s
                                , "SELECT SESSION_PLAN FROM ROLES WHERE ID = 'TEAMBONUSCONFIG'"
                                , null
                                , SerializerReadDouble.INSTANCE).find();
                
            } catch (BasicException e) {
    		l.info("BasicException e:" + e);
    		//implement later
            }
    	for (int i = 0; i < rows; i++) {


    	    Object[] rowcontent = new Object [4];
            try {
                rowcontent = (Object[]) new PreparedSentence(s
                    , "SELECT ID, BONUS_FROM_SALES, SESSION_PLAN,  TEAMBONUS_COEF FROM ROLES WHERE ROWNUMBER = ? AND ID != 'TEAMBONUSCONFIG'"
                    , SerializerWriteInteger.INSTANCE
                    , new SerializerReadBasic(new Datas[] {Datas.STRING, Datas.DOUBLE, Datas.DOUBLE, Datas.DOUBLE })).find(i+1);
                } catch (BasicException e) {
                    l.info("BasicException e: " + e);
                }
            if (rowcontent != null) {
                mapIndivBonus.put(rowcontent[0].toString(), (double)rowcontent[1]);
                mapIndivPlan.put(rowcontent[0].toString(), (double)rowcontent[2]);
                mapTeamBonusCoef.put(rowcontent[0].toString(), (double)rowcontent[3]);
            }
            
 

    	}
        
        
    }

    public final boolean isTeamBonusConfigHost(String hostname) throws BasicException {
        return 0 <= (Integer) new PreparedSentence(s
            , "SELECT COUNT(HOST) FROM TEAMBONUSCONFIGHOSTS WHERE HOST = ?"
            , SerializerWriteString.INSTANCE
            , SerializerReadInteger.INSTANCE).find(hostname);
    }

    
    public int countRolesRows() {
        int r = 0;
        try {
            r = (int) new PreparedSentence (s
    				, "SELECT COUNT(ID) FROM ROLES"
        			, null
        			, SerializerReadInteger.INSTANCE).find();
            } catch (BasicException e) {
    		l.info("BasicException e:" + e);
    		//implement later
            }
        return r;
    }
    
}

