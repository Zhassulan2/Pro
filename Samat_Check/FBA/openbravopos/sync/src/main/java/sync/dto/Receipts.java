package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "receipts", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Receipts implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "money")
  private String money;
  @Column(name = "datenew")
  private java.sql.Timestamp datenew;
  @Column(name = "attributes")
  private String attributes;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getMoney() {
    return money;
  }

  public void setMoney(String money) {
    this.money = money;
  }


  public java.sql.Timestamp getDatenew() {
    return datenew;
  }

  public void setDatenew(java.sql.Timestamp datenew) {
    this.datenew = datenew;
  }


  public String getAttributes() {
    return attributes;
  }

  public void setAttributes(String attributes) {
    this.attributes = attributes;
  }


  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }

  public String getIdentifier() {return this.id;}

  public String getIdentifierName(){return "id";}

}
