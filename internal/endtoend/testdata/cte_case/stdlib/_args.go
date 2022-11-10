[]compiler.paramRef{
        
    compiler.paramRef{
        parent:(*ast.A_Expr){
            Kind:ast.A_Expr_Kind(1),
            Name:(*ast.List){
                Items:[]ast.Node{
                    (*ast.String){Str:"="}
                }
            }, 
            Lexpr:(*ast.ColumnRef){
                Name:"",
                Fields:(*ast.List){
                    Items:[]ast.Node{
                        (*ast.String){Str:"label"}
                    }
                },
                Location:229
            },
            Rexpr:(*ast.ParamRef){Number:1, Location:237, Dollar:true},
            Location:235},
            rv:(*ast.RangeVar)(nil),
            ref:(*ast.ParamRef){Number:1,Location:237,Dollar:true},
            name:"",
        },


    compiler.paramRef{
        parent:(*ast.A_Expr){
            Kind:ast.A_Expr_Kind(1),
            Name:(*ast.List){
                Items:[]ast.Node{
                    (*ast.String){Str:"="}
                }
            },
            Lexpr:(*ast.ColumnRef){
                Name:"", Fields:(*ast.List){
                    Items:[]ast.Node{
                        (*ast.String){Str:"ready"}
                    }
                },
                Location:245
            },
            Rexpr:(*ast.ParamRef){Number:2, Location:253, Dollar:true},
            Location:251},
            rv:(*ast.RangeVar)(nil),
            ref:(*ast.ParamRef){Number:2, Location:253, Dollar:true},
            name:""
        }
}