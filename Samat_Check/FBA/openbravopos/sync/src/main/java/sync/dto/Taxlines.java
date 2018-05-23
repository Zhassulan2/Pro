package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "taxlines", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Taxlines implements DataModel{
  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "receipt")
  private String receipt;
  @Column(name = "taxid")
  private String taxid;
  @Column(name = "base")
  private double base;
  @Column(name = "amount")
  private double amount;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getReceipt() {
    return receipt;
  }

  public void setReceipt(String receipt) {
    this.receipt = receipt;
  }


  public String getTaxid() {
    return taxid;
  }

  public void setTaxid(String taxid) {
    this.taxid = taxid;
  }


  public double getBase() {
    return base;
  }

  public void setBase(double base) {
    this.base = base;
  }


  public double getAmount() {
    return amount;
  }

  public void setAmount(double amount) {
    this.amount = amount;
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
