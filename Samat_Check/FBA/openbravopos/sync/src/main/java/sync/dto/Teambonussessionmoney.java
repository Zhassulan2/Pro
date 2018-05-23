package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "teambonussessionmoney", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Teambonussessionmoney implements DataModel{
    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getTbs_id() {
        return tbs_id;
    }

    public void setTbs_id(String tbs_id) {
        this.tbs_id = tbs_id;
    }

    public String getMoney() {
        return money;
    }

    public void setMoney(String money) {
        this.money = money;
    }

    @Id
    @Column(name="Id",unique = true, nullable = false)
    private String id;

    @Column(name = "Tbs_id")
    private String tbs_id;

    @Column(name = "Money")
    private String money;


    public String getIdentifier() {return this.id;}

    public String getIdentifierName(){return "id";}

}
