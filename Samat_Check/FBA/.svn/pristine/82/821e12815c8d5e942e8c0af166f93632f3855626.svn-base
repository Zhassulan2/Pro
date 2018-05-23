/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package net.ddns.evcalyptus.bonus;

import com.openbravo.basic.BasicException;
import com.openbravo.data.loader.*;
import com.openbravo.pos.forms.AppUser;
import com.openbravo.pos.forms.BeanFactoryDataSingle;
import com.openbravo.pos.payment.PaymentInfo;
import com.openbravo.pos.ticket.TicketInfo;
import com.openbravo.pos.util.ThumbNailBuilder;
import java.util.List;
import java.util.UUID;
import java.util.logging.Logger;

import javax.swing.ImageIcon;
import com.openbravo.pos.ticket.UserInfo;

/**
 *
 * @author user2
 */
public class DataLogicIndivBonus extends BeanFactoryDataSingle{
	Logger l = Logger.getLogger(getClass().getName());

    public DataLogicIndivBonus() {
    }
    
    protected Session s;
    
    static final double BONUS_FROM_SALES = 0.01;
        
    //protected DataLogicBonusConfig dl_BonusConfig;
    
    protected SentenceFind m_bonuscountperson;
    protected SentenceList m_bonuslistperson;
    protected SentenceFind m_bonusgetperson;
    protected SentenceFind m_ismoneypersoninserted;
    protected SentenceExec m_bonusinsertperson;
    protected SentenceFind m_bonussumsales;
    protected SentenceFind m_bonussumbonus;
    protected SerializerRead peopleread;
    protected SentenceExec bonusinsertticket;
    protected SentenceExec bonusdeleteticket;
    protected SentenceFind getrole;
    
    
    //private DataLogicBonusConfig dl_BonusConfig;
    
    public void init(Session s){
        
        //dl_BonusConfig = new DataLogicBonusConfig();
    	//dl_BonusConfig = (DataLogicBonusConfig) m_App.getBean("net.ddns.evcalyptus.bonus.DataLogicBonusConfig");
    	
    	
        final ThumbNailBuilder tnb = new ThumbNailBuilder(32, 32, "com/openbravo/images/yast_sysadmin.png");        
        peopleread = new SerializerRead() {
            public Object readValues(DataRead dr) throws BasicException {
                return new AppUser(
                        dr.getString(1),
                        dr.getString(2),
                        dr.getString(3),
                        dr.getString(4),
                        dr.getString(5),
                        new ImageIcon(tnb.getThumbNail(ImageUtils.readImage(dr.getBytes(6)))));                
            }
        };
        
        
        
        
        
        
                // bonus logic count number of unique persons in active cache.
        m_bonuscountperson = new StaticSentence(s
        	, "SELECT DISTINCT COUNT(PERSON) FROM BONUSTICKETS WHERE MONEY = ?"
        	, SerializerWriteString.INSTANCE
                , SerializerReadInteger.INSTANCE);
                
        m_bonuslistperson = new StaticSentence(s
                , "SELECT DISTINCT B.PERSON, P.NAME, P.APPPASSWORD, P.CARD, P.ROLE, P.IMAGE FROM PEOPLE P, BONUSTICKETS B WHERE B.PERSON = P.ID AND B.MONEY = ?"
                , SerializerWriteString.INSTANCE
                , peopleread);
        
        m_ismoneypersoninserted = new StaticSentence(s
                , "SELECT COUNT(ID) FROM BONUSHISTORY WHERE MONEY = ? AND PERSON = ?"
                , new SerializerWriteBasic(Datas.STRING, Datas.STRING)
                , SerializerReadInteger.INSTANCE);
                        
        m_bonusgetperson = new StaticSentence(s
        	, "SELECT MONEY, PERSON, PAYMENTTOTAL, BONUS FROM BONUSTICKETS WHERE MONEY = ? AND PERSON = ?"
                //need to sum PAYMENTTOTAL AND BONUS
        	, new SerializerWriteBasic(new Datas[] {Datas.STRING, Datas.STRING})
        	, new SerializerReadBasic(new Datas[] {Datas.STRING, Datas.STRING, Datas.DOUBLE, Datas.DOUBLE}));
        
        m_bonusinsertperson = new StaticSentence(s
                , "INSERT INTO BONUSHISTORY(ID, MONEY, PERSON, DAILYSALES, DAILYBONUS) " +
                  "VALUES (?, ?, ?, ?, ?)"
                , new SerializerWriteBasic(new Datas[] {Datas.STRING, Datas.STRING, Datas.STRING, Datas.DOUBLE, Datas.DOUBLE}));
        
        m_bonussumsales = new StaticSentence(s
                , "SELECT SUM(PAYMENTTOTAL) FROM BONUSTICKETS WHERE MONEY = ? AND PERSON = ?" 
                ,  new SerializerWriteBasic(new Datas[] {Datas.STRING, Datas.STRING})
                ,  SerializerReadDouble.INSTANCE);
       
        m_bonussumbonus = new StaticSentence(s
                , "SELECT SUM(BONUS) FROM BONUSTICKETS WHERE MONEY = ? AND PERSON = ?" 
                ,  new SerializerWriteBasic(new Datas[] {Datas.STRING, Datas.STRING})
                ,  SerializerReadDouble.INSTANCE);
        
        getrole =  new StaticSentence (s
                , "SELECT ROLE FROM PEOPLE WHERE ID = ?"
                , SerializerWriteString.INSTANCE
                , SerializerReadString.INSTANCE);
        
        bonusinsertticket = new PreparedSentence(s
    	        , "INSERT INTO BONUSTICKETS (TICKET, MONEY, PERSON, PAYMENTTOTAL, BONUS) VALUES (?, ?, ?, ?, ?)"
    	        , SerializerWriteParams.INSTANCE);
        
        bonusdeleteticket =  new StaticSentence(s
                    , "DELETE FROM BONUSTICKETS WHERE TICKET = ?"
                    , SerializerWriteString.INSTANCE);
         
    }
    
    
    public final int bonusCountPerson(String sActiveCashIndex) throws BasicException {
    Integer i = (Integer) m_bonuscountperson.find(sActiveCashIndex);
    return i.intValue();
    } 
    
    public final List bonusListPerson(String sActiveCashIndex) throws BasicException {
        return m_bonuslistperson.list(sActiveCashIndex);
    }      
    
    public final int isMoneyPersonInserted (String money, String person) throws BasicException {
        return (Integer) m_ismoneypersoninserted.find(money, person);
    }
    
    public final int isMoneyPersonInserted (Object[] money_person) throws BasicException {
        return isMoneyPersonInserted(money_person[0].toString(), money_person[1].toString());
    }
    
    public final Object[] bonusGetPerson(Object[] money_person) throws BasicException {
        //Datas [] bonusdatas = new Datas[] {Datas.STRING, Datas.STRING, Datas.DOUBLE, Datas.DOUBLE};
        //bonusdatas = m_bonusgetperson.find(sPersonID);
        return (Object[]) m_bonusgetperson.find(money_person);
    } 
    
    public final void bonusInsertPerson(Object[] bonus, double sumpayments, double sumbonus) throws BasicException {
        //int updateresult = ((Object[]) params)[5] == null
             
        Object[] bonusUpdated = new Object[] {
            UUID.randomUUID().toString(),
            bonus [0],
            bonus [1],
            sumpayments, //bonus [2],
            sumbonus //bonus [3]
        };
        m_bonusinsertperson.exec(bonusUpdated);
    }
    
    public final double bonusSumPayments(Object[] money_person) throws BasicException {
        double sumpayments = (Double) m_bonussumsales.find(money_person);
        return sumpayments;//i.intValue();
    }
    
    public final double bonusSumBonus(Object[] money_person) throws BasicException {
        double sumbonus = (Double) m_bonussumbonus.find(money_person);
        return sumbonus;//i.intValue();
    }
     
    public String getRole (String userID) throws BasicException {
    
   
                        String role = getrole.find(userID).toString();
    			return role; 
    }
    
    public void bonusAddTicket (final TicketInfo ticket, double indivbonusrate) throws BasicException {
         // update bonustickets
 

                String role = getRole(ticket.getUser().getId().toString());

                for (final PaymentInfo b : ticket.getPayments()) {
                        bonusinsertticket.exec(new DataParams() { public void writeValues() throws BasicException {
                        	setString(1, ticket.getId());
                        	setString(2, ticket.getActiveCash());
                            setString(3, ticket.getUser().getId());
                            setDouble(4, b.getTotal());
                            setDouble(5, b.getTotal() * indivbonusrate);
                        }});
     
                }
     }
     
    public void bonusDeleteTicket (final TicketInfo ticket) throws BasicException {
        bonusdeleteticket.exec(ticket.getId());
    }
}


