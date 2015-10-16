from flask import Flask, render_template, request, Response, jsonify
from deliverables import postHandler, getDeliverablesForModule, deleteDeliverableHandler

import logging

app = Flask(__name__)
# Note: We don't need to call run() since our application is embedded within
# the App Engine WSGI application server.

@app.route('/')
def hello():
  return render_template('index.html')

@app.route('/admin')
def admin():
  """Return the admin page"""
  return render_template('admin.html')


@app.route('/admin/api/module/<moduleName>', methods=["GET"])
def module_api_handler(moduleName):
  results = getDeliverablesForModule(request, moduleName)
  logging.info("Results:\n{}".format(results))
  import json
  return Response(json.dumps(results),  mimetype='application/json')

@app.route('/admin/api/module/<moduleName>/deliverable/<title>', methods=["POST", "GET", "DELETE"])
def deliverable_api_handler(moduleName, title):
  """REST API Handler for deliverable api"""
  if request.method == "POST":
    return postHandler(request, moduleName, title)
  if request.method == "DELETE":
    return deleteDeliverableHandler(request, moduleName, title)

@app.route("/admin/api/module/<moduleName>/deliverable/<title>/getUploadURL", methods=["GET"])
def getUploadURL():
  uploadUrl = blobstore.create_upload_url('/admin/api/module/{moduleName}/deliverable/{title}/submit'.format(moduleName=moduleName, title=title))
  return Response(json.dumps(uploadUri), mimetype="application/json")

@app.route("/admin/api/module/<moduleName>/deliverable/<title>/submit", methods=["POST"])
def fileSubmission(moduleName, title):
  return fileSubmissionHandler(request, moduleName, title)

@app.errorhandler(404)
def page_not_found(e):
  """Return a custom 404 error."""
  return render_template('404.html'), 404


@app.errorhandler(500)
def application_error(e):
  """Return a custom 500 error."""
  return render_template('500.html'), 500
