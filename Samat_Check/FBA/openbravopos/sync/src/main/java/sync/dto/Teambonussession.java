package sync.dto;


import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;
import java.sql.Timestamp;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "teambonussession", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Teambonussession implements DataModel{
    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public int getTbsequence() {
        return tbsequence;
    }

    public void setTbsequence(int tbsequence) {
        this.tbsequence = tbsequence;
    }

    public Timestamp getDatestart() {
        return datestart;
    }

    public void setDatestart(Timestamp datestart) {
        this.datestart = datestart;
    }

    public Timestamp getDateend() {
        return dateend;
    }

    public void setDateend(Timestamp dateend) {
        this.dateend = dateend;
    }

    public float getTeambonussales() {
        return teambonussales;
    }

    public void setTeambonussales(float teambonussales) {
        this.teambonussales = teambonussales;
    }

    public float getTeambonus() {
        return teambonus;
    }

    public void setTeambonus(float teambonus) {
        this.teambonus = teambonus;
    }

    @Id
    @Column(name="id",unique = true, nullable = false)
    private String id;
    @Column(name = "tbsequence")
    private int tbsequence;
    @Column(name = "datestart")
    private java.sql.Timestamp datestart;
    @Column(name = "dateend")
    private java.sql.Timestamp dateend;
    @Column(name = "teambonussales")
    private float teambonussales;
    @Column(name = "teambonus")
    private float teambonus;

    public String getIdentifier() {return this.id;}

    public String getIdentifierName(){return "id";}
}
