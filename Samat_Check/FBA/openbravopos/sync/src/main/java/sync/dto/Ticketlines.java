package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "ticketlines")
public class Ticketlines implements DataModel{

  /*@Id
  @Column(name = "ticket")
  private String ticket;

  @Column(name = "line")
  private int line;*/

    @EmbeddedId
    TicketLine id;

  @Column(name = "product")
  private String product;
  @Column(name = "attributesetinstance_Id")
  private String attributesetinstance_Id;
  @Column(name = "units")
  private double units;
  @Column(name = "price")
  private double price;
  @Column(name = "taxid")
  private String taxid;
  @Column(name = "attributes")
  private String attributes;
  @Column(name = "sync")
  private long sync;


  public TicketLine getId() {
    return id;
  }

  public void setId(TicketLine id) {
    this.id = id;
  }

  public String getProduct() {
    return product;
  }

  public void setProduct(String product) {
    this.product = product;
  }


  public String getAttributesetinstance_Id() {
    return attributesetinstance_Id;
  }

  public void setAttributesetinstance_Id(String attributesetinstance_Id) {
    this.attributesetinstance_Id = attributesetinstance_Id;
  }


  public double getUnits() {
    return units;
  }

  public void setUnits(double units) {
    this.units = units;
  }


  public double getPrice() {
    return price;
  }

  public void setPrice(double price) {
    this.price = price;
  }


  public String getTaxid() {
    return taxid;
  }

  public void setTaxid(String taxid) {
    this.taxid = taxid;
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

  public String getIdentifier() {return this.id.getTicket();}

  public String getIdentifierName(){return "ticket";}

}
