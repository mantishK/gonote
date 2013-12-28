$(document).ready(noteInit)
function noteInit() {
	getNotes()
	addNote()
	editNote()
}
function editNote() {
	$("#note-edit-btn").click(editNoteClick);
}
function editNoteClick() {
	var title = $("#edit-note-title").val()
	var content = $("#edit-note-content").val()
	var note_id = $("#edit-note-id").val()
	$.ajax({
	  type: "PUT",
	  url: "api/note/",
	  data: {title:title,content:content,note_id:note_id}
	}).done(postEdit);
}
function postEdit(data) {
	$('#edit-note-modal').modal('hide')
	var alertHtml = ''
	if(data.response == "ok") {
		alertHtml = '<div class="alert alert-success" data-dismiss="alert" id="alert-message">Note updated successfuly</div>'
	}
	else {
		alertHtml = '<div class="alert alert-danger" data-dismiss="alert" id="alert-error">Error while updating note</div>'
	}
	$('#alert-wrapper').html(alertHtml)
	$('#note-list').html('')
	getNotes()
}
function editDeleteNote() {
	$(".note").mouseenter(showEditDeleteBtn);
	$(".note").mouseleave(hideEditDeleteBtn);
	$(".edit-btn span").click(editAction);
	$(".delete-btn span").click(deleteAction);
}
function editAction(evt) {
	var noteId = $(this).parent().parent().attr("note_id");
	$.ajax({
	  type: "GET",
	  url: "api/note/details?note_id="+noteId
	}).done(postDetails);
}
function postDetails(data) {
	if(data.response == "ok") {
		$('#edit-note-modal').modal('show')
		$('#edit-note-title').val(data.note.Title)
		$('#edit-note-content').val(data.note.Content)
		$('#edit-note-id').val(data.note.Note_id)
	}
	else {
		alertHtml = '<div class="alert alert-danger" data-dismiss="alert" id="alert-error">Error while retriving note</div>'
	}
	
}
function deleteAction(evt) {
	var noteId = $(this).parent().parent().attr("note_id");
	$.ajax({
	  type: "DELETE",
	  url: "api/note/?note_id="+noteId
	}).done(postDelete);
}
function postDelete(data) {
	var alertHtml = ''
	if(data.response == "ok") {
		alertHtml = '<div class="alert alert-success" data-dismiss="alert" id="alert-message">Note deleted successfuly</div>'
	}
	else {
		alertHtml = '<div class="alert alert-danger" data-dismiss="alert" id="alert-error">Error while deleting note</div>'
	}
	$('#alert-wrapper').html(alertHtml)
	$('#note-list').html('')
	getNotes()
}
function showEditDeleteBtn(evt) {
	$(this).find(".delete-btn").css('visibility','visible')
	$(this).find(".edit-btn").css('visibility','visible')
}
function hideEditDeleteBtn(evt) {
	$(this).find(".delete-btn").css('visibility','hidden')
	$(this).find(".edit-btn").css('visibility','hidden')
}
function addNote() {
	$("#note-save-btn").click(addNoteClick);
}
function addNoteClick() {
	var title = $("#new-note-title").val()
	var content = $("#new-note-content").val()
	$.post("api/note/",{title:title,content:content},postAdd)
}
function postAdd(data) {
	var title = $("#new-note-title").val('')
	var content = $("#new-note-content").val('')
	$('#new-note-modal').modal('hide')
	var alertHtml = ''
	if(data.response == "ok") {
		alertHtml = '<div class="alert alert-success" data-dismiss="alert" id="alert-message">Note added successfuly</div>'
	}
	else {
		alertHtml = '<div class="alert alert-danger" data-dismiss="alert" id="alert-error">Error while adding note</div>'
	}
	$('#alert-wrapper').html(alertHtml)
	$('#note-list').html('')
	getNotes()	
	
}
function getNotes() {
	$.getJSON("api/note/",{},displayNotes);
}
function displayNotes(data) {
	var htmlContent = ""
	for (var i = 0; i < data.notes.length; i++) {
		if(i%4==0) {
			if(i != 0)
				htmlContent += '</div>'
			htmlContent += '<div class="row">'
		}
		htmlContent += '<div class="col-md-3"> <div class="note">'
		htmlContent += '<div class="title">'+data.notes[i].Title+'</div>'
		htmlContent += '<div class="content"><p>'+data.notes[i].Content+'</p></div>'
		htmlContent += '<div class="row" note_id="' + data.notes[i].Note_id +'">'+
		'<div class="col-md-3 delete-btn" ><span class="glyphicon glyphicon-remove"></span></div>'+
		'<div class="col-md-3 edit-btn"><span class="glyphicon glyphicon-edit"></span></div>'
		htmlContent += '<div class="col-md-5 col-md-offset-1 date-wrapper"><div class="date">'+ (new Date(data.notes[i].Modified * 1000)).toDateString()+'</div></div>'
		htmlContent += '</div> </div></div>'	
	}
	htmlContent += '</div>'
	$("#note-list").append(htmlContent)
	editDeleteNote()
}
