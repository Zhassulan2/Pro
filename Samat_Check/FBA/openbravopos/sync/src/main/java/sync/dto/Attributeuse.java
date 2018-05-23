package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "attributeuse", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Attributeuse implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "attributeset_Id")
  private String attributeset_Id;
  @Column(name = "attribute_Id")
  private String attribute_Id;
  @Column(name = "lineno")
  private long lineno;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getAttributeset_Id() {
    return attributeset_Id;
  }

  public void setAttributeset_Id(String attributeset_Id) {
    this.attributeset_Id = attributeset_Id;
  }


  public String getAttribute_Id() {
    return attribute_Id;
  }

  public void setAttribute_Id(String attribute_Id) {
    this.attribute_Id = attribute_Id;
  }


  public long getLineno() {
    return lineno;
  }

  public void setLineno(long lineno) {
    this.lineno = lineno;
  }


  public long getSync() {
    return sync;
  }

  public void setSync(long sync) {
    this.sync = sync;
  }

  public String getIdentifier() {
    return this.id;
  }
  public String getIdentifierName(){return "id";}
}
