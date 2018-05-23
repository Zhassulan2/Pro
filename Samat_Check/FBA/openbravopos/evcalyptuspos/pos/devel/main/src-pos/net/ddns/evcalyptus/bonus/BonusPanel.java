/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package net.ddns.evcalyptus.bonus;

import javax.swing.ListCellRenderer;
import com.openbravo.data.gui.ListCellRendererBasic;
import com.openbravo.data.loader.ComparatorCreator;
import com.openbravo.data.loader.TableDefinition;
import com.openbravo.data.loader.Vectorer;
import com.openbravo.data.user.EditorRecord;
import com.openbravo.data.user.ListProvider;
import com.openbravo.data.user.ListProviderCreator;
import com.openbravo.data.user.SaveProvider;
import com.openbravo.pos.admin.DataLogicAdmin;
//import com.openbravo.pos.admin.RolesView;
import com.openbravo.pos.forms.AppLocal;
import com.openbravo.pos.panels.JPanelTable;


/**
 *
 * @author adrianromero
 */
/**
 *
 * @author user2
 */
public class BonusPanel extends JPanelTable {
    
    private TableDefinition troles;
    private BonusView jeditor;
    
    /** Creates a new instance of BonusPanel */
    public BonusPanel() {
     }
    
    protected void init() {
        DataLogicAdmin dlAdmin  = (DataLogicAdmin) app.getBean("com.openbravo.pos.admin.DataLogicAdmin");
 
        troles = dlAdmin.getTableRoles();                 
        jeditor = new BonusView(dirty, app);
    }
    
    public ListProvider getListProvider() {
        return new ListProviderCreator(troles);
    }
    
    public SaveProvider getSaveProvider() {
        return new SaveProvider(troles);        
    }
    
    public Vectorer getVectorer() {
        return troles.getVectorerBasic(new int[] {5});
    }
    
    public ComparatorCreator getComparatorCreator() {
        return troles.getComparatorCreator(new int[] {5});
    }
    
    public ListCellRenderer getListCellRenderer() {
        return new ListCellRendererBasic(troles.getRenderStringBasic(new int[] {5}));
    }
    
    public EditorRecord getEditor() {
        return jeditor;
    }
    
    public String getTitle() {
        return AppLocal.getIntString("Menu.RolesBonus");
    }        
}


