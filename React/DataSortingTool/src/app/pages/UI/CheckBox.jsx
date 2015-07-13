var React = require('react');
var Bootstrap = require('bootstrap');

/*
 * Class: CheckBox
 * Usage:
 *      -To create available  data types for the user to select and sort
 * Design:
 *      -Does Not contain text or labels for the user to see what the checkbox is for, so you need to create text/label/p/header/... within the UI container
 *          to inform the user the context of the checkbox allowing free placement and more flexible UIs for the designer
 *      -Requires a callback function called checkDataType, so that the event change within this class can pass information back to the parent node (ex. UI container)
 *          on the status of the checkbox
 *      -Var checked is the return type for the callback method checkDataType
 *      -HandleClick is the event function for onClick to handle the value of checked
*/

var CheckBox = React.createClass({
    getInitialState: function(){
        return{checked: true}
    },
    propTypes: {
        checkDataType: React.PropTypes.func.isRequired
    },
    //checkDataType(checkBoxID, checked)
    //checkBoxID should be the prop.ID of the class, which should be named after the data type you are trying to select (ex. gender)
    //checked is the true/false var that determines if the prop selected should be removed from or added to the array: active[] in DataTypeSelection
    handleClick: function(event){
        this.setState({
            checked: !this.state.checked
        });
        console.log("CheckBox Id: " + this.props.id + " Check Box Status: " + this.state.checked);
        this.props.checkDataType(this.props.id, this.state.checked);
    },
    render: function(){
        return(
            <input name={this.props.id} id={this.props.id} type="checkbox" onClick={this.handleClick}></input>
        );
    }
});

module.exports = CheckBox;
