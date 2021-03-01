# AppEncryptionDeclarationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppEncryptionDeclaration**](AppEncryptionDeclaration.md) |  | 
**Included** | Pointer to [**[]App**](App.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppEncryptionDeclarationResponse

`func NewAppEncryptionDeclarationResponse(data AppEncryptionDeclaration, links DocumentLinks, ) *AppEncryptionDeclarationResponse`

NewAppEncryptionDeclarationResponse instantiates a new AppEncryptionDeclarationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppEncryptionDeclarationResponseWithDefaults

`func NewAppEncryptionDeclarationResponseWithDefaults() *AppEncryptionDeclarationResponse`

NewAppEncryptionDeclarationResponseWithDefaults instantiates a new AppEncryptionDeclarationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppEncryptionDeclarationResponse) GetData() AppEncryptionDeclaration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppEncryptionDeclarationResponse) GetDataOk() (*AppEncryptionDeclaration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppEncryptionDeclarationResponse) SetData(v AppEncryptionDeclaration)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppEncryptionDeclarationResponse) GetIncluded() []App`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppEncryptionDeclarationResponse) GetIncludedOk() (*[]App, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppEncryptionDeclarationResponse) SetIncluded(v []App)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppEncryptionDeclarationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppEncryptionDeclarationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppEncryptionDeclarationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppEncryptionDeclarationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


