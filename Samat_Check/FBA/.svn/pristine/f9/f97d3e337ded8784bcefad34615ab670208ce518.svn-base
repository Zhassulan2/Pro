/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.openbravo.pos.reports;

import com.openbravo.basic.BasicException;
import com.openbravo.data.loader.Datas;
import com.openbravo.data.loader.QBFCompareEnum;
import com.openbravo.data.loader.SerializerWrite;
import com.openbravo.data.loader.SerializerWriteBasic;
import com.openbravo.pos.forms.AppView;
import com.openbravo.pos.forms.AppUser;
//import com.openbravo.pos.ticket.UserInfo;
import java.awt.Component;

/**
 *
 * @author user2
 */
public class JParamsCurrentUser extends javax.swing.JPanel implements ReportEditorCreator{

   

        
    private AppUser currentuser;
    private AppView app;
    
    
    public JParamsCurrentUser() {
        //initComponents();
    }

    
    @Override
    public Object createValue() throws BasicException {
        return new Object[] {QBFCompareEnum.COMP_EQUALS, currentuser.getId()};
    }

    @Override
    public SerializerWrite getSerializerWrite() {
        return new SerializerWriteBasic(new Datas[] {Datas.OBJECT, Datas.STRING});
    }


    @Override
    public void init(AppView app) {
        //currentuser = app.getAppUserView().getUser();
        this.app = app;
    }

    @Override
    public void activate() throws BasicException {
        //currentuser = null;
        currentuser = app.getAppUserView().getUser();
    }

    @Override
    public Component getComponent() {
        return this;
    }
    

    
}
