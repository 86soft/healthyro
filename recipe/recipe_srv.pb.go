// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: recipe/recipe_srv.proto

package recipe

import (
	common "github.com/86soft/healthyro/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_recipe_recipe_srv_proto protoreflect.FileDescriptor

var file_recipe_recipe_srv_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x1a, 0x17, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x5f, 0x6d, 0x73, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa6, 0x05, 0x0a, 0x0d, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x22, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72,
	0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x12, 0x25, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70,
	0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12,
	0x25, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f, 0x2e, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x72, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x38, 0x36, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x72, 0x6f,
	0x2f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_recipe_recipe_srv_proto_goTypes = []interface{}{
	(*ListRecipesRequest)(nil),              // 0: healthyro.recipe.ListRecipesRequest
	(*GetRecipeRequest)(nil),                // 1: healthyro.recipe.GetRecipeRequest
	(*CreateRecipeRequest)(nil),             // 2: healthyro.recipe.CreateRecipeRequest
	(*UpdateRecipeTitleRequest)(nil),        // 3: healthyro.recipe.UpdateRecipeTitleRequest
	(*UpdateRecipeDescriptionRequest)(nil),  // 4: healthyro.recipe.UpdateRecipeDescriptionRequest
	(*UpdateRecipeExternalLinkRequest)(nil), // 5: healthyro.recipe.UpdateRecipeExternalLinkRequest
	(*DeleteRecipeRequest)(nil),             // 6: healthyro.recipe.DeleteRecipeRequest
	(*ListRecipesResponse)(nil),             // 7: healthyro.recipe.ListRecipesResponse
	(*GetRecipeResponse)(nil),               // 8: healthyro.recipe.GetRecipeResponse
	(*CreateRecipeResponse)(nil),            // 9: healthyro.recipe.CreateRecipeResponse
	(*common.Empty)(nil),                    // 10: healthyro.common.Empty
}
var file_recipe_recipe_srv_proto_depIdxs = []int32{
	0,  // 0: healthyro.recipe.RecipeService.ListRecipes:input_type -> healthyro.recipe.ListRecipesRequest
	1,  // 1: healthyro.recipe.RecipeService.GetRecipe:input_type -> healthyro.recipe.GetRecipeRequest
	2,  // 2: healthyro.recipe.RecipeService.CreateRecipe:input_type -> healthyro.recipe.CreateRecipeRequest
	3,  // 3: healthyro.recipe.RecipeService.UpdateRecipeTitle:input_type -> healthyro.recipe.UpdateRecipeTitleRequest
	4,  // 4: healthyro.recipe.RecipeService.UpdateRecipeDescription:input_type -> healthyro.recipe.UpdateRecipeDescriptionRequest
	5,  // 5: healthyro.recipe.RecipeService.UpdateRecipeExternalLink:input_type -> healthyro.recipe.UpdateRecipeExternalLinkRequest
	6,  // 6: healthyro.recipe.RecipeService.DeleteRecipe:input_type -> healthyro.recipe.DeleteRecipeRequest
	7,  // 7: healthyro.recipe.RecipeService.ListRecipes:output_type -> healthyro.recipe.ListRecipesResponse
	8,  // 8: healthyro.recipe.RecipeService.GetRecipe:output_type -> healthyro.recipe.GetRecipeResponse
	9,  // 9: healthyro.recipe.RecipeService.CreateRecipe:output_type -> healthyro.recipe.CreateRecipeResponse
	10, // 10: healthyro.recipe.RecipeService.UpdateRecipeTitle:output_type -> healthyro.common.Empty
	10, // 11: healthyro.recipe.RecipeService.UpdateRecipeDescription:output_type -> healthyro.common.Empty
	10, // 12: healthyro.recipe.RecipeService.UpdateRecipeExternalLink:output_type -> healthyro.common.Empty
	10, // 13: healthyro.recipe.RecipeService.DeleteRecipe:output_type -> healthyro.common.Empty
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_recipe_recipe_srv_proto_init() }
func file_recipe_recipe_srv_proto_init() {
	if File_recipe_recipe_srv_proto != nil {
		return
	}
	file_recipe_recipe_msg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_recipe_recipe_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recipe_recipe_srv_proto_goTypes,
		DependencyIndexes: file_recipe_recipe_srv_proto_depIdxs,
	}.Build()
	File_recipe_recipe_srv_proto = out.File
	file_recipe_recipe_srv_proto_rawDesc = nil
	file_recipe_recipe_srv_proto_goTypes = nil
	file_recipe_recipe_srv_proto_depIdxs = nil
}