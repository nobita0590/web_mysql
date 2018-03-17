"use strict";function __export(e){for(var o in e)exports.hasOwnProperty(o)||(exports[o]=e[o])}var __decorate=this&&this.__decorate||function(e,o,a,t){var n,r=arguments.length,l=r<3?o:null===t?t=Object.getOwnPropertyDescriptor(o,a):t;if("object"==typeof Reflect&&"function"==typeof Reflect.decorate)l=Reflect.decorate(e,o,a,t);else for(var _=e.length-1;_>=0;_--)(n=e[_])&&(l=(r<3?n(l):r>3?n(o,a,l):n(o,a))||l);return r>3&&l&&Object.defineProperty(o,a,l),l},core_1=require("@angular/core"),common_1=require("@angular/common"),forms_1=require("@angular/forms"),table_component_1=require("./components/table.component");exports.DataTable=table_component_1.DataTable;var column_component_1=require("./components/column.component");exports.DataTableColumn=column_component_1.DataTableColumn;var row_component_1=require("./components/row.component");exports.DataTableRow=row_component_1.DataTableRow;var pagination_component_1=require("./components/pagination.component");exports.DataTablePagination=pagination_component_1.DataTablePagination;var header_component_1=require("./components/header.component");exports.DataTableHeader=header_component_1.DataTableHeader;var px_1=require("./utils/px"),hide_1=require("./utils/hide"),min_1=require("./utils/min"),default_translations_type_1=require("./types/default-translations.type");exports.defaultTranslations=default_translations_type_1.defaultTranslations,__export(require("./tools/data-table-resource")),exports.DATA_TABLE_DIRECTIVES=[table_component_1.DataTable,column_component_1.DataTableColumn];var DataTableModule=function(){function e(){}return e}();DataTableModule=__decorate([core_1.NgModule({imports:[common_1.CommonModule,forms_1.FormsModule],declarations:[table_component_1.DataTable,column_component_1.DataTableColumn,row_component_1.DataTableRow,pagination_component_1.DataTablePagination,header_component_1.DataTableHeader,px_1.PixelConverter,hide_1.Hide,min_1.MinPipe],exports:[table_component_1.DataTable,column_component_1.DataTableColumn]})],DataTableModule),exports.DataTableModule=DataTableModule;