package sync.dto;

import com.google.gson.annotations.SerializedName;
import org.hibernate.annotations.OptimisticLockType;

import javax.persistence.*;

@Entity
@org.hibernate.annotations.Entity(optimisticLock = OptimisticLockType.ALL, dynamicUpdate = true)
@Table(name = "products", uniqueConstraints = {@UniqueConstraint(columnNames = "id")})
public class Products implements DataModel{
  @Id
  @Column(name = "id", unique = true, nullable = false)
  @SerializedName("Id")
  private String id;
  @Column(name = "reference")
  @SerializedName("Reference")
  private String reference;
  @Column(name = "code")
  @SerializedName("Code")
  private String code;
  @Column(name = "codetype")
  @SerializedName("Codetype")
  private String codetype;
  @Column(name = "name")
  @SerializedName("Name")
  private String name;
  @Column(name = "pricebuy")
  @SerializedName("Pricebuy")
  private java.lang.Double pricebuy;
  @Column(name = "pricesell")
  @SerializedName("Pricesell")
  private java.lang.Double pricesell;
  @Column(name = "category")
  @SerializedName("Category")
  private String category;
  @Column(name = "taxcat")
  @SerializedName("Taxcat")
  private String taxcat;
  @Column(name = "attributeset_Id")
  @SerializedName("Attributeset_Id")
  private String attributeset_Id;
  @Column(name = "stockcost")
  @SerializedName("Stockcost")
  private java.lang.Double stockcost;
  @Column(name = "stockvolume")
  @SerializedName("Stockvolume")
  private java.lang.Double stockvolume;
  @Column(name = "image")
  @SerializedName("Image")
  private String image;
  @Column(name = "iscom")
  @SerializedName("Iscom")
  private String iscom;
  @Column(name = "isscale")
  @SerializedName("Isscale")
  private String isscale;
  @Column(name = "attributes")
  @SerializedName("Attributes")
  private String attributes;
  @Column(name = "sync")
  private long sync;


  public String getId() {
    return id;
  }

  public void setId(String id) {
    this.id = id;
  }


  public String getReference() {
    return reference;
  }

  public void setReference(String reference) {
    this.reference = reference;
  }


  public String getCode() {
    return code;
  }

  public void setCode(String code) {
    this.code = code;
  }


  public String getCodetype() {
    return codetype;
  }

  public void setCodetype(String codetype) {
    this.codetype = codetype;
  }


  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }


  public double getPricebuy() {
    return pricebuy;
  }

  public void setPricebuy(java.lang.Double pricebuy) {
    this.pricebuy = pricebuy;
  }


  public double getPricesell() {
    return pricesell;
  }

  public void setPricesell(java.lang.Double pricesell) {
    this.pricesell = pricesell;
  }


  public String getCategory() {
    return category;
  }

  public void setCategory(String category) {
    this.category = category;
  }


  public String getTaxcat() {
    return taxcat;
  }

  public void setTaxcat(String taxcat) {
    this.taxcat = taxcat;
  }


  public String getAttributeset_Id() {
    return attributeset_Id;
  }

  public void setAttributeset_Id(String attributeset_Id) {
    this.attributeset_Id = attributeset_Id;
  }


  public double getStockcost() {
    return stockcost;
  }

  public void setStockcost(java.lang.Double stockcost) {
    this.stockcost = stockcost;
  }


  public double getStockvolume() {
    return stockvolume;
  }

  public void setStockvolume(java.lang.Double stockvolume) {
    this.stockvolume = stockvolume;
  }


  public String getImage() {
    return image;
  }

  public void setImage(String image) {
    this.image = image;
  }


  public String getIscom() {
    return iscom;
  }

  public void setIscom(String iscom) {
    this.iscom = iscom;
  }


  public String getIsscale() {
    return isscale;
  }

  public void setIsscale(String isscale) {
    this.isscale = isscale;
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

  public String getIdentifier() {return this.id;}

  public String getIdentifierName(){return "id";}

}
