(function($) {
"use strict";

var ProductForm = React.createClass({displayName: "ProductForm",
    handleChange: function() {
    },
    render: function() {
        return (
            React.createElement("form", {className: "ProductForm"}, 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement(Input, {placeholder: "商品コード", maxLength: 20, required: true})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement(Input, {placeholder: "商品名", maxLength: 30, required: false})
                )
            )
        );
    }
});

var OrderForm = React.createClass({displayName: "OrderForm",
    render: function() {
        return (
            React.createElement("form", {className: "OrderForm"}, 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("label", null, "注文日時"), 
                    React.createElement(DateTimeInput, {placeholder: "注文日時"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("label", null, "発送指示日"), 
                    React.createElement(DateInput, {placeholder: "発送指示日"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("select", null, 
                        React.createElement("option", null, "未入金"), React.createElement("option", null, "入金済み")
                    )
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("input", {type: "text", className: "form-control", placeholder: "郵便番号"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("input", {type: "text", className: "form-control", placeholder: "お届け先住所"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("input", {type: "text", className: "form-control", placeholder: "受取人電話番号"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("input", {type: "text", className: "form-control", placeholder: "キース商品コード"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("input", {type: "text", className: "form-control", placeholder: "商品名"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("input", {type: "number", className: "form-control", placeholder: "個数"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement(DateTimeInput, {placeholder: "配送時間"})
                ), 
                React.createElement("div", {className: "form-group"}, 
                    React.createElement("textarea", {className: "form-control", placeholder: "備考"})
                )
            )
        );
    }
});

var Page = React.createClass({displayName: "Page",
    getInitialState: function() {
        return {
            orderFormVisible: false,
            productFormVisible: false,
            text: ''
        };
    },
    handleClick: function(descriptor) {
        switch (descriptor) {
            case 'OrderForm':
                this.setState({ orderFormVisible: true });
                break;
            case 'ProductForm':
                this.setState({ productFormVisible: true });
                break;
        }
    },
    closeDialog: function(descriptor) {
        switch (descriptor) {
            case 'OrderForm':
                this.setState({ orderFormVisible: false });
                break;
            case 'ProductForm':
                this.setState({ productFormVisible: false });
                break;
        }
    },
    handleCancelClick: function(descriptor) {
        if (!confirm('ダイアログを閉じます。編集内容が失われますが、よろしいですか？')) return;
        this.closeDialog(descriptor);
    },
    handleOkClick: function(descriptor) {
        this.closeDialog(descriptor);
    },
    render: function() {
        return (
            React.createElement("div", {className: "Page"}, 
                React.createElement("button", {className: "btn btn-primary", "data-toggle": "modal", onClick: this.handleClick.bind(this,'OrderForm')}, 
                    "オーダーフォーム表示"
                ), 
                React.createElement("button", {className: "btn btn-primary", "data-toggle": "modal", onClick: this.handleClick.bind(this,'ProductForm')}, 
                    "商品フォーム表示"
                ), 
                React.createElement(ModalDialog, {visible: this.state.orderFormVisible, 
                             descriptor: "OrderForm", 
                             title: "注文", 
                             cancelCaption: "閉じる", 
                             okCaption: "保存", 
                             onCancelClick: this.handleCancelClick.bind(this,'OrderForm'), 
                             onOkClick: this.handleOkClick.bind(this,'OrderForm'), 
                             content: React.createElement(OrderForm, null)}
                ), 
                React.createElement(ModalDialog, {visible: this.state.productFormVisible, 
                             descriptor: "ProductForm", 
                             title: "商品", 
                             cancelCaption: "閉じる", 
                             okCaption: "保存", 
                             onCancelClick: this.handleCancelClick.bind(this,'ProductForm'), 
                             onOkClick: this.handleOkClick.bind(this,'ProductForm'), 
                             content: React.createElement(ProductForm, null)}
                )
            )
        );
    }
});

React.render(
    React.createElement(Page, null),
    document.getElementById('page')
);

})(jQuery);
