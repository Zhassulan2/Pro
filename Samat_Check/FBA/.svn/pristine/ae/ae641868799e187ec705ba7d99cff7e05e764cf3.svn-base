package sync.dto;

/*
CREATE TABLE public.sync_deleted_records (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	tablename varchar(500) NOT NULL,
	pk varchar(500) NOT NULL
)
WITH (
	OIDS=FALSE
) ;
* */

import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "sync_deleted_records", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})

public class Syncdeletedrecords implements DataModel{
    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getTablename() {
        return tablename;
    }

    public void setTablename(String tablename) {
        this.tablename = tablename;
    }

    public String getPk() {
        return pk;
    }

    public void setPk(String pk) {
        this.pk = pk;
    }

    @Id
    @Column(name = "id", unique = true, nullable = false)
    private String id;
    @Column(name = "tablename")
    private String tablename;
    @Column(name = "pk")
    private String pk;

    public String getIdentifier() {
        return this.id;
    }
    public String getIdentifierName(){return "id";}
}
