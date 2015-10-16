# deliverables.py
from flask import make_response, abort
from google.appengine.ext import ndb
from google.appengine.ext import blobstore
from google.appengine.ext.webapp import blobstore_handlers
import logging


class Deliverable(ndb.Model):
  title = ndb.StringProperty()
  file_key = ndb.BlobKeyProperty()
  date = ndb.DateTimeProperty(auto_now_add=True)

def moduleKey(moduleName=None):
  """Constructs a Datastore key for a deliverable entity with moduleName."""
  return ndb.Key('Deliverable', moduleName or 'default_module')

def deliverableKey(moduleName, title):
  return ndb.Key('Deliverable', title, parent=moduleKey(moduleName))

def postHandler(request, moduleName, title):
  #If no key provided, create a new deliverable
  d = Deliverable(id=title, parent=moduleKey(moduleName))
  d.title = title
  #TODO: Do something for the filekey
  d.put()
  logging.info("Saved a deliverable with key: {}".format(d.key.urlsafe()))
  return make_response("")


def deleteDeliverableHandler(request, moduleName, title):
  k = deliverableKey(moduleName, title)
  if not k.get():
    return make_response("No such entity with title: {} and moduleName:{}".format(title, moduleName)), 403
  k.delete()  
  return make_response("")


def getDeliverablesForModule(request, moduleName):
  results = Deliverable.query(ancestor=moduleKey(moduleName)).order(Deliverable.date).fetch()
  return [{"key": r.key.urlsafe(), "title": r.title, "moduleName": moduleName} for r in results ]


def fileSubmissionHandler(request, moduleName, title):
  blob_key = get_blob_key(request, "file")
  d = deliverableKey(moduleName, title).get()
  #TODO: May be need to delete the old file, check in local datastore later.
  d.file_key = blob_key
  d.put()
  return make_response("")

def get_blob_key(request, field_name):
  """
  Parse out and return the blob key from a file uploaded via the App Engine Blobstore API.
  """

  upload_file = request.files[field_name]
  header = upload_file.headers['Content-Type']
  parsed_header = parse_options_header(header)
  blob_key = parsed_header[1]['blob-key']
  return blob_key