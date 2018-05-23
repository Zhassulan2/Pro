package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "closedcash")
public class Closedcash implements DataModel{

  @Id
  @Column(name = "money")
  private String money;
  @Column(name = "host")
  private String host;
  @Column(name = "hostsequence")
  private long hostsequence;
  @Column(name = "datestart")
  private java.sql.Timestamp datestart;
  @Column(name = "dateend")
  private java.sql.Timestamp dateend;
  @Column(name = "sync")
  private long sync;


  public String getMoney() {
    return money;
  }

  public void setMoney(String money) {
    this.money = money;
  }


  public String getHost() {
    return host;
  }

  public void setHost(String host) {
    this.host = host;
  }


  public long getHostsequence() {
    return hostsequence;
  }

  public void setHostsequence(long hostsequence) {
    this.hostsequence = hostsequence;
  }


  public java.sql.Timestamp getDatestart() {
    return datestart;
  }

  public void setDatestart(java.sql.Timestamp datestart) {
    this.datestart = datestart;
  }


  public java.sql.Timestamp getDateend() {
    return dateend;
  }

  public void setDateend(java.sql.Timestamp dateend) {
    this.dateend = dateend;
  }


  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }

  public String getIdentifier() {
    return this.money;
  }
  public String getIdentifierName(){return "money";}
}
