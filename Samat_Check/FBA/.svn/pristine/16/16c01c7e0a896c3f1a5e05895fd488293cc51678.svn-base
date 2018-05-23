package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "places", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Places implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;
  @Column(name = "name")
  private String name;
  @Column(name = "X")
  private long X;
  @Column(name = "Y")
  private long Y;
  @Column(name = "floor")
  private String floor;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }


  public long getX() {
    return X;
  }

  public void setX(long X) {
    this.X = X;
  }


  public long getY() {
    return Y;
  }

  public void setY(long Y) {
    this.Y = Y;
  }


  public String getFloor() {
    return floor;
  }

  public void setFloor(String floor) {
    this.floor = floor;
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
