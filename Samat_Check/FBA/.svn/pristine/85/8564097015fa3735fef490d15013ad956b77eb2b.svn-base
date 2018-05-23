/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

    /*  working notes and planning.
    The idea is to create a new type of session, which includes several Activecash/MONEY sessions of individual hosts connected to central DB (super MONEY session).  the tbsession is created either the first time at client startup
    or like MONEY at JPanelCloseMoney the moment previous tbsession is closed.  tbsession represents a working day (or other period depending when Cash is closed) with several connected hosts, which contains the conbined sales (and teambonuses) 
    for these combined clients.  TEAMBONUSCONFIGHOSTS is needed to specify hosts, which participate into tbsession, so as to exclude non-selling hosts, which may never see their cash closed.
    
    At JRootApp startup: 
    = public final boolean isTeamBonusConfigHosts(String hostname) - check if the current host is eligible for teambonus. if not, ignore TeamBonus logic.  If yes, proceed with TeamBonus logic 
    = public final Object[] getTeamBonusSession() - get the current session (if there is one) from TEAMBONUSSESSION, i.e. one with DATEEND = null; if there is none, call setTeamBonusSession() and insertTeamBonusSession(Object[] tbsession)
    = public final boolean isTeamBonusActiveSessionMoney(String money) - check if the open TEAMBONUSSESSION already contains current money=m_App.getActiveCashIndex(), if yes just proceed, if no, do insertTeamBonusSessionMoney;
    = public void insertTeamBonusSessionMoney(String tbsession, String money) - insert current m_App.getActiveCashIndex() into TEAMBONUSSESSIONMONEY;
    = public void setTeamBonusSession(String tbs_id, Date dStart, Date dEnd) - create new session (if no existing found) with given parameters;
    = public void insertTeamBonusSession(Object[] tbsession) - insert newly created session into TEAMBONUSSESSION and remember to call from JRootApp insertTeamBonusSessionMoney(tb_Session.getTeamBonusActiveSessionId(), m_App.getActiveCashIndex());
    
    
    at JPanelCloseMoney:
    = public final boolean checkTeamBonusSessionLastMoney() - "SELECT COUNT(C.MONEY) FROM CLOSEDCASH C, TEAMBONUSSESSIONMONEY T WHERE C.MONEY = T.MONEY AND C.DATEEND IS NULL" check if all MONEY in the current TEAMBONUSSESSION have been closed DATEEND <> 0.  If 
    = if not all MONEY except current host have been closed, close MONEY and add "SELECT SUM(DAILYSALES) FROM BONUSHISTORY WHERE MONEY = ? GROUP BY MONEY" to TEAMBONUSSESSION.TEAMBONUSSALES.
    = public void closeTeamBonusSession() - if this is the last MONEY in TEAMBONUSSESSIONMONEY, close TEAMBONUSSESSIONMONEY and close MONEY.  add "SELECT SUM(DAILYSALES) FROM BONUSHISTORY WHERE MONEY = ? GROUP BY MONEY" to TEAMBONUSSESSION.TEAMBONUSSALES.  
    = updateTeamBonus - The last client calculates all TEAMBONUSSESSION.TEAMBONUS and TEAMBONUSSESSIONPERSON.TEAMBONUSFORPERSON and inserts accordingly
    = setTeamBonusSession() + insertTeamBonusSession() + insertTeamBonusSessionMoney(), just like MONEY (or m_App.getActiveCashIndex()) new session is created at the moment the last session is closed.
    */

package net.ddns.evcalyptus.bonus;

import com.openbravo.basic.BasicException;
import com.openbravo.data.loader.*;
import com.openbravo.pos.forms.AppUser;
import com.openbravo.pos.forms.BeanFactoryDataSingle;
import com.openbravo.pos.util.ThumbNailBuilder;

import java.util.Date;
import java.util.List;
import java.util.UUID;
import java.util.logging.Logger;

import javax.swing.ImageIcon;

/**
 *
 * @author user2
 */
public class DataLogicTeamBonus extends BeanFactoryDataSingle{
    //static final double TB_FROM_SALES = 0.01;
	
    static final Logger l = Logger.getLogger(DataLogicTeamBonus.class.getName());
    
    private String tb_SessionId;
    private int tb_SessionSeq;
    private Date tb_SessionDateStart;
    private Date tb_SessionDateEnd;
    private double tb_SessionSales;
    private double tb_SessionBonus;
    
    
    //protected SerializerRead peopleread;
    final ThumbNailBuilder tnb = new ThumbNailBuilder(32, 32, "com/openbravo/images/yast_sysadmin.png");
    protected SerializerRead peopleread = new SerializerRead() {
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
    
    
    protected SentenceFind tb_confighosts;
    protected SentenceFind tb_activesession;
    protected SentenceFind tb_checkactivesessionmoney;
    protected SentenceExec tb_inserttbsessionmoney;
    protected SentenceExec tb_inserttbsession;
    protected SentenceFind tb_sessionseq;
    protected SentenceFind tb_countactivemoney;
    protected SentenceFind tb_selectactivemoney;
    protected SentenceExec tb_updatesessionnumbers;
    protected SentenceFind tb_getactivemoneysales;
    protected SentenceFind tb_getsessionbonus;
    protected SentenceList tb_listperson;
    protected SentenceExec tb_insertforperson;
    protected SentenceFind tb_issessionplanok;
    protected SentenceExec tb_settbsessionbonuszero;
    protected SentenceExec tb_closesession;
    protected SentenceFind tb_isnomoneyofhostinactivetbsession;
    
    
    //protected Datas[] tbsessiondatas;
    
    public DataLogicTeamBonus() {
    }
    
    
    @Override
    public void init(Session s) {
        
        //tbsessiondatas = new Datas[] {Datas.STRING, Datas.TIMESTAMP, Datas.TIMESTAMP, Datas.DOUBLE, Datas.DOUBLE};

        tb_confighosts = new PreparedSentence(s
            , "SELECT COUNT(HOST) FROM TEAMBONUSCONFIGHOSTS WHERE HOST = ?"
            , SerializerWriteString.INSTANCE
            , SerializerReadInteger.INSTANCE);
            
        tb_activesession = new StaticSentence(s
            , "SELECT ID, TBSEQUENCE, DATESTART, DATEEND FROM TEAMBONUSSESSION WHERE DATEEND IS NULL"
            , null
            , new SerializerReadBasic(new Datas[] {Datas.STRING, Datas.INT, Datas.TIMESTAMP, Datas.TIMESTAMP}));
        
        tb_checkactivesessionmoney = new StaticSentence(s
            , "SELECT COUNT(TBSM.MONEY) FROM TEAMBONUSSESSIONMONEY TBSM WHERE TBSM.MONEY = ?"
            , SerializerWriteString.INSTANCE
            , SerializerReadInteger.INSTANCE);
        
        tb_inserttbsessionmoney = new StaticSentence(s
            , "INSERT INTO TEAMBONUSSESSIONMONEY(ID, TBS_ID, MONEY) " +
              "VALUES (?, ?, ?)"
            , new SerializerWriteBasic(new Datas[] {Datas.STRING, Datas.STRING, Datas.STRING}));
        
        tb_inserttbsession = new StaticSentence(s
            , "INSERT INTO TEAMBONUSSESSION(ID, TBSEQUENCE, DATESTART, TEAMBONUSSALES, TEAMBONUS) " +
              "VALUES (?, ?, ?, ?, ?)"
            , new SerializerWriteBasic(new Datas[] {Datas.STRING, Datas.INT, Datas.TIMESTAMP, Datas.DOUBLE, Datas.DOUBLE}));
        
        tb_sessionseq = new StaticSentence(s
            ,"SELECT MAX(TBSEQUENCE) FROM TEAMBONUSSESSION"
            , null
            , SerializerReadInteger.INSTANCE);
        
        tb_countactivemoney = new StaticSentence(s
            , "SELECT COUNT(C.MONEY) FROM CLOSEDCASH C, TEAMBONUSSESSIONMONEY T WHERE C.MONEY = T.MONEY AND C.DATEEND IS NULL"
            , null
            , SerializerReadInteger.INSTANCE);
        
        tb_selectactivemoney = new StaticSentence(s
                , "SELECT C.MONEY FROM CLOSEDCASH C, TEAMBONUSSESSIONMONEY T WHERE C.MONEY = T.MONEY AND C.DATEEND IS NULL"
                , null
                , SerializerReadString.INSTANCE);
        
        tb_getactivemoneysales = new StaticSentence(s
                , "SELECT SUM(DAILYSALES) FROM BONUSHISTORY WHERE MONEY = ? GROUP BY MONEY"
                , SerializerWriteString.INSTANCE
                , SerializerReadDouble.INSTANCE);
        
        tb_updatesessionnumbers = new StaticSentence(s
                , "UPDATE TEAMBONUSSESSION SET TEAMBONUSSALES = (TEAMBONUSSALES + ?), TEAMBONUS = (TEAMBONUS + ?) WHERE ID = ?"
                , new SerializerWriteBasic(new Datas[] {Datas.DOUBLE, Datas.DOUBLE, Datas.STRING})
                );
        
        tb_getsessionbonus = new StaticSentence(s
                , "SELECT SUM(TEAMBONUS) FROM TEAMBONUSSESSION WHERE ID = ? GROUP BY ID"
                , SerializerWriteString.INSTANCE
                , SerializerReadDouble.INSTANCE);
        
        tb_listperson = new StaticSentence(s
                , "SELECT DISTINCT B.PERSON, P.NAME, P.APPPASSWORD, P.CARD, P.ROLE, P.IMAGE FROM PEOPLE P, BONUSHISTORY B, TEAMBONUSSESSIONMONEY TBSM WHERE B.PERSON = P.ID AND B.MONEY = TBSM.MONEY AND TBSM.TBS_ID = ?"
                , SerializerWriteString.INSTANCE
                , peopleread);
        
        tb_insertforperson = new StaticSentence(s
        	, "INSERT INTO TEAMBONUSSESSIONPERSON (ID, TBS_ID, PERSON, TEAMBONUSFORPERSON)" +
        	  "VALUES (?, ?, ?, ?)"
        	, new SerializerWriteBasic(new Datas[] {Datas.STRING, Datas.STRING, Datas.STRING, Datas.DOUBLE})
                );
        
        tb_issessionplanok = new StaticSentence(s
                , "SELECT TEAMBONUSSALES FROM TEAMBONUSSESSION WHERE ID = ?"
                , SerializerWriteString.INSTANCE
                , SerializerReadDouble.INSTANCE);
        
        tb_settbsessionbonuszero = new StaticSentence(s
                , "UPDATE TEAMBONUSSESSION SET TEAMBONUS = 0.0 WHERE ID = ?"
                , SerializerWriteString.INSTANCE);
        
        tb_closesession = new StaticSentence(s
                , "UPDATE TEAMBONUSSESSION SET DATEEND = ? WHERE ID = ?"
                , new SerializerWriteBasic(new Datas[] {Datas.TIMESTAMP, Datas.STRING})
                );
        
        tb_isnomoneyofhostinactivetbsession = new StaticSentence(s
        		, "SELECT COUNT(TBSM.MONEY) FROM TEAMBONUSSESSIONMONEY TBSM, TEAMBONUSSESSION TBS, CLOSEDCASH C " + 
        				"WHERE TBSM.TBS_ID = TBS.ID AND TBSM.MONEY = C.MONEY AND TBS.DATEEND IS NULL AND C.HOST = ?"
        		, SerializerWriteString.INSTANCE
        		, SerializerReadInteger.INSTANCE);
        
    }
    
    public final boolean isTeamBonusConfigHost(String hostname) {
        int i = 0;
        try {
            	i = (Integer) tb_confighosts.find(hostname);
            	//l.info("hostname = " + hostname);
        } catch (BasicException e) {
            	l.info("BasicException e: " + e);
            	//TODO : implement MessageInf msg = new MessageInf(MessageInf.SGN_NOTICE, AppLocal.getIntString("message.cannotinsertteambonus"), e);
            	//msg.show(this);
            } 
        return i != 0;
    }
    
    public final Object[] getTeamBonusActiveSession() throws BasicException {
        Object[] tbsession = new Object[3];
        tbsession = (Object[]) tb_activesession.find();
        if (tbsession != null){
            tb_SessionId = tbsession[0].toString();
            tb_SessionSeq = (int)tbsession[1];
            tb_SessionDateStart = (Date)tbsession[2];
        }
        return tbsession;
    }
    
    public final boolean isTeamBonusActiveSessionMoney(String money) throws BasicException {
        int i = (Integer)tb_checkactivesessionmoney.find(money);
        return i != 0;
        /*if (i == 0) 
            return false; 
        else 
            return true;
        */
    }
    
    public void insertTeamBonusSessionMoney(String tbsessionid, String money) throws BasicException {
    	//need to insert check for already existing money in the tbsession in order to avoid same money in various tbsessons
           	tb_inserttbsessionmoney.exec(UUID.randomUUID().toString(), tbsessionid, money);
   }
    
    public void insertTeamBonusActiveSessionMoney(String money) throws BasicException {
    	insertTeamBonusSessionMoney(tb_SessionId, money);
    }
    
    public void setTeamBonusActiveSession(String tbs_id, int tbsseq, Date dStart) throws BasicException{
        tb_SessionId = tbs_id;
        tb_SessionSeq = tbsseq;
        tb_SessionDateStart = dStart;
        try {
        	insertTeamBonusActiveSession();
        } catch (BasicException e) {
        	l.info("BasicException e: " + e);
        	//TODO : implement MessageInf msg = new MessageInf(MessageInf.SGN_NOTICE, AppLocal.getIntString("message.cannotinsertteambonus"), e);
            //msg.show(this);
        }
    } 
    
    public void setTeamBonusActiveSession(String tbs_id, Date dStart) throws BasicException{
        tb_SessionId = tbs_id;    	
    	tb_SessionSeq = tb_sessionseq.find() == null ? 0 : (Integer)tb_sessionseq.find()+ 1;
        tb_SessionDateStart = dStart;
        setTeamBonusActiveSession(tb_SessionId, tb_SessionSeq, tb_SessionDateStart);
    }
    
    public void setTeamBonusActiveSession(String tbs_id) throws BasicException{
    	tb_SessionId = tbs_id;    	
    	tb_SessionSeq = tb_sessionseq.find() == null ? 0 : (Integer)tb_sessionseq.find()+ 1;
        tb_SessionDateStart = new Date();
        setTeamBonusActiveSession(tb_SessionId, tb_SessionSeq, tb_SessionDateStart);
    }
    
    public void setTeamBonusActiveSession() throws BasicException{
        tb_SessionId = UUID.randomUUID().toString();
        tb_SessionSeq = tb_sessionseq.find() == null ? 0 : (Integer)tb_sessionseq.find()+ 1;
        tb_SessionDateStart = new Date();
        setTeamBonusActiveSession(tb_SessionId, tb_SessionSeq, tb_SessionDateStart);
    }
    
    public void setTeamBonusActiveSession(Object[] tbsession) throws BasicException{
        setTeamBonusActiveSession(tbsession[0].toString(), (int)tbsession[1], (Date)tbsession[2]);
    }
    
    public String getTeamBonusActiveSessionId(){
        return tb_SessionId;
    }
    
    public int getTeamBonusActiveSessionSeq(){
        return tb_SessionSeq;
    }
    
    public Date getTeamBonusActiveSessionDateStart(){
        return tb_SessionDateStart;
    }
    
    public void insertTeamBonusActiveSession() throws BasicException {
        if (tb_SessionId != null & tb_SessionDateStart != null) {
            try {
                insertTeamBonusSession(new Object[] {tb_SessionId, tb_SessionSeq, tb_SessionDateStart, 0.0, 0.0});
            } catch (BasicException e) {
        	l.info("BasicException e: " + e);
        	//TODO : implement MessageInf msg = new MessageInf(MessageInf.SGN_NOTICE, AppLocal.getIntString("message.cannotinsertteambonus"), e);
            //msg.show(this);
            }
                    
        }
        else l.info("Cannot insert session since no session has been set or loaded");
    }
        
    public void insertTeamBonusSession(Object[] tbsession) throws BasicException {
        try {
            tb_inserttbsession.exec(tbsession);
        } catch (BasicException e) {
        	l.info("BasicException e: " + e);
        	//TODO : implement MessageInf msg = new MessageInf(MessageInf.SGN_NOTICE, AppLocal.getIntString("message.cannotinsertteambonus"), e);
            //msg.show(this);
            }
    }
    
    public final boolean isTeamBonusSessionLastMoney(String money) throws BasicException{
        boolean ret = false;
        try {
        	int i = 0;
            i = (Integer)tb_countactivemoney.find();
            if (i == 1) {
                String selectActiveMoney = tb_selectactivemoney.find().toString();
                if (money.equals(selectActiveMoney))
                    ret = true;
                else {
                    l.info("ActiveMoney: " + money + " does not coincide with the only remaining unclosed money in DB: "
                            + selectActiveMoney);
                    ret = false;
                }
            }
            else ret = false;
        } catch (BasicException e) {
        	l.info("BasicException e: " + e);
        	//TODO : implement MessageInf msg = new MessageInf(MessageInf.SGN_NOTICE, AppLocal.getIntString("message.cannotinsertteambonus"), e);
            //msg.show(this);
        }
        return ret;
    }
    
    public void updateTeamBonusSessionNumbers(String money, double tb_from_sales) throws BasicException {
        double sessionsales = (double)tb_getactivemoneysales.find(money);
        double sessionbonus = sessionsales * tb_from_sales;
        tb_updatesessionnumbers.exec(new Object[] {sessionsales, sessionbonus, tb_SessionId});
    }    
    
    public boolean isTeamPlanOK (String tbsessionid, double tb_session_plan) throws BasicException {
        double tbsessionsales = (Double) tb_issessionplanok.find(tbsessionid);
        return tbsessionsales >= tb_session_plan; 
    }
    
    public void setTeamBonusToZero (String tbsessionid) throws BasicException {
        tb_settbsessionbonuszero.exec(tbsessionid);
    }
    
    public void setTeamBonusActiveSessionToZero() throws BasicException {
        //tb_settbsessionbonuszero.exec(this.getTeamBonusActiveSessionId());
        setTeamBonusToZero(this.getTeamBonusActiveSessionId());
    }
    
    public void closeTeamBonusActiveSession() throws BasicException {
        if (tb_SessionId != null & tb_SessionDateStart != null) {
            tb_SessionDateEnd = new Date();
            closeTeamBonusSession(new Object[] {tb_SessionDateEnd, tb_SessionId});
            tb_SessionDateEnd = null;
        }
        else l.info("Cannot close session since no session has been set or loaded");
    }
        
    public void closeTeamBonusSession(Object[] tbclosesession) throws BasicException {
    	try {
    		java.util.List teambonuspeople = tb_listperson.list(tb_SessionId);
    		for (int i = 0; i < teambonuspeople.size(); i++) {
    			AppUser user = (AppUser) teambonuspeople.get(i);
    			double totalsessionbonus = tb_getsessionbonus.find(tb_SessionId) == null 
    					? 0 
    					: (double)tb_getsessionbonus.find(tb_SessionId);
                        //DataLogicBonusConfig.mapTeamBonusCoef.get(user.getRole());
                        double teambonusforpers = (DataLogicBonusConfig.mapTeamBonusCoef.get(user.getRole()) * totalsessionbonus) / teambonuspeople.size();
    			tb_insertforperson.exec(new Object[] {UUID.randomUUID().toString(), tb_SessionId, user.getId(), teambonusforpers});
    		}
    		tb_closesession.exec(tbclosesession);
    	} catch (BasicException e) {
        	l.info("BasicException e: " + e);
        	//TODO : implement MessageInf msg = new MessageInf(MessageInf.SGN_NOTICE, AppLocal.getIntString("message.cannotinsertteambonus"), e);
            //msg.show(this);
        }
    }

public final boolean isNoMoneyOfHostInActiveTBSession(String hostname){
    	int i = 0;
		try {
			i= (Integer)tb_isnomoneyofhostinactivetbsession.find(hostname);
    	}catch (BasicException e) {
        	l.info("BasicException e: " + e);
        	//TODO : implement MessageInf msg = new MessageInf(MessageInf.SGN_NOTICE, AppLocal.getIntString("message.cannotinsertteambonus"), e);
            //msg.show(this);
        }
		return i == 0;
    }
 
    
}
