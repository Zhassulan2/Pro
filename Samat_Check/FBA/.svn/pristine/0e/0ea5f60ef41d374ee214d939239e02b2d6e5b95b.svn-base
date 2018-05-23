package sync.dto;

import org.hibernate.annotations.OptimisticLockType;
import javax.persistence.*;
import java.io.UnsupportedEncodingException;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "payments", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Payments implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "receipt")
  private String receipt;
  @Column(name = "payment")
  private String payment;
  @Column(name = "total")
  private double total;
  @Column(name = "transid")
  private String transid;
  @Column(name = "returnmsg")
  private String returnmsg;
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


  public String getPayment() {
    return payment;
  }

  public void setPayment(String payment) {
    this.payment = payment;
  }


  public double getTotal() {
    return total;
  }

  public void setTotal(double total) {
    this.total = total;
  }


  public String getTransid() {
    return transid;
  }

  public void setTransid(String transid) {
    this.transid = transid;
  }


  public String getReturnmsg() throws UnsupportedEncodingException {

      return returnmsg;
  }

  public void setReturnmsg(String returnmsg) {

    this.returnmsg = returnmsg;
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
