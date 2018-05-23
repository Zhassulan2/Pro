package sync.dto;

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "roles", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Roles implements DataModel{

  @Id
  @Column(name = "id", unique = true, nullable = false)
  private String id;

  @Column(name = "rownumber")
  private int rownumber;

  public int getRownumber() {
    return rownumber;
  }

  public void setRownumber(int rownumber) {
    this.rownumber = rownumber;
  }

  public Double getBonus_from_sales() {
    return bonus_from_sales;
  }

  public void setBonus_from_sales(Double bonus_from_sales) {
    this.bonus_from_sales = bonus_from_sales;
  }

  public Double getSession_plan() {
    return session_plan;
  }

  public void setSession_plan(Double session_plan) {
    this.session_plan = session_plan;
  }

  public Double getTeambonus_coef() {
    return teambonus_coef;
  }

  public void setTeambonus_coef(Double teambonus_coef) {
    this.teambonus_coef = teambonus_coef;
  }

  @Column(name = "bonus_from_sales")
  private Double bonus_from_sales;

  @Column(name = "session_plan")
  private Double session_plan;

  @Column(name = "teambonus_coef")
  private Double teambonus_coef;

  @Column(name = "name")
  private String name;
  @Column(name = "permissions")
  private String permissions;
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


  public String getPermissions() {
    return permissions;
  }

  public void setPermissions(String permissions) {
    this.permissions = permissions;
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
