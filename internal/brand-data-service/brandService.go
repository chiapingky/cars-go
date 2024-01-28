package branddataservice

import "log"

type BrandServiceImpl struct {
	BrandAccessor
}

func (brandService BrandServiceImpl) GetBrandByName(name string) Brand {
	brand, err := brandService.BrandAccessor.FindByName(name)
	if err != nil {
		log.Default().Println("GetBrandByName() | Error when getting brand ", err)
		return Brand{}
	}
	return brand
}

func (brandService BrandServiceImpl) InsertBrand(brand Brand) Brand {
	checkedBrand := brandService.GetBrandByName(brand.Name)
	if (checkedBrand != Brand{}) {
		log.Default().Println("InsertBrand() | Brand already exist ")
		return Brand{}
	}
	saveResult, err := brandService.BrandAccessor.Save(brand)
	if err != nil {
		log.Default().Println("InsertBrand() | Error when saving brand ", err)
		return Brand{}
	}
	return saveResult
}

func (brandService BrandServiceImpl) DeleteBrand(brand Brand) Brand {
	checkedBrand := brandService.GetBrandByName(brand.Name)
	if (checkedBrand == Brand{}) {
		log.Default().Println("DeleteBrand() | Brand doesn't exist ")
		return Brand{}
	}
	deleteResult, err := brandService.BrandAccessor.Delete(brand)
	if err != nil {
		log.Default().Println("DeleteBrand() | Error when deleting brand ", err)
		return Brand{}
	}
	return deleteResult
}
