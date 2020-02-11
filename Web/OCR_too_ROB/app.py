from flask import Flask, render_template, request
from werkzeug import secure_filename
import os
from subprocess import PIPE, Popen
import sys
from PIL import Image
import pytesseract
import argparse
import cv2
import os
app = Flask(__name__)
UPLOAD_FOLDER = '/tmp'
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER
app.config['MAX_CONTENT_LENGTH'] = 10 * 1024 * 1024


@app.route("/")
def index():
  return render_template("index.html", info="", error="")

@app.route('/validate', methods = ['GET', 'POST'])
def upload_file():
   if request.method == 'POST':
      f = request.files['file']

      # create a secure filename
      filename = secure_filename(f.filename)

      # save file to /static/uploads
      filepath = os.path.join(app.config['UPLOAD_FOLDER'],filename)
      print(filepath)
      f.save(filepath)

      pipe = Popen(pytesseract.pytesseract.tesseract_cmd + ' ' + filepath + ' stdout', shell=True, stdout=PIPE, stderr=PIPE)
      output = str(pipe.stdout.read()).replace("\n", "").replace("\r", "")
      p = Popen('echo "' + output + '" | grep allowed', shell=True, stdout=PIPE, stderr=PIPE)
      stdout, stderr = p.communicate()


      return render_template("index.html", info=output, error=str(stderr))

if __name__ == '__main__':
   app.run(host="0.0.0.0", port=5000, debug=True)
