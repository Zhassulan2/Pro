package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "tickets", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Tickets implements DataModel{
  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "tickettype")
  private long tickettype;
  @Column(name = "ticketid")
  private long ticketid;
  @Column(name = "person")
  private String person;
  @Column(name = "customer")
  private String customer;
  @Column(name = "status")
  private long status;
 //@Column(name = "sync")
 // private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public long getTickettype() {
    return tickettype;
  }

  public void setTickettype(long tickettype) {
    this.tickettype = tickettype;
  }


  public long getTicketid() {
    return ticketid;
  }

  public void setTicketid(long ticketid) {
    this.ticketid = ticketid;
  }


  public String getPerson() {
    return person;
  }

  public void setPerson(String person) {
    this.person = person;
  }


  public String getCustomer() {
    return customer;
  }

  public void setCustomer(String customer) {
    this.customer = customer;
  }


  public long getStatus() {
    return status;
  }

  public void setStatus(long status) {
    this.status = status;
  }

/*
  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }
*/
  public String getIdentifier() {return this.id;}

  public String getIdentifierName(){return "id";}

}
