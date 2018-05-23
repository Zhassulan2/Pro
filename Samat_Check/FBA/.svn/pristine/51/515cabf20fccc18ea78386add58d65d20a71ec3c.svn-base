package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "attributevalue", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Attributevalue implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "attribute_Id")
  private String attribute_Id;
  @Column(name = "value")
  private String value;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getAttribute_Id() {
    return attribute_Id;
  }

  public void setAttribute_Id(String attribute_Id) {
    this.attribute_Id = attribute_Id;
  }


  public String getValue() {
    return value;
  }

  public void setValue(String value) {
    this.value = value;
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
