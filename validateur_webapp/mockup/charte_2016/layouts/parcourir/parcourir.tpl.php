<div class="<?php print $classes ?>" <?php if (!empty($css_id)) { print "id=\"$css_id\""; } ?>>
  <?php if ($content['navigation']): ?>
    <div class="panel-navigation">
      <div class="container">
        <div class="row row-navigation">
          <?php if ($content['navigation']): ?>
            <?php print $content['navigation']; ?>
          <?php endif; ?>
        </div>
      </div>
    </div>
  <?php endif; ?>
  <div class="panel-main">
  <?php if ($content['block2'] || $content['block3']): ?>
    <div class="row">
      <?php if ($content['block2']): ?>
        <div class="col-md-8">
          <?php print $content['block2']; ?>
        </div>
      <?php endif; ?>
      <?php if ($content['block3']): ?>
        <div class="col-md-4">
          <?php print $content['block3']; ?>
        </div>
      <?php endif; ?>
    </div>
  <?php endif; ?>

  <?php if ($content['block4'] || $content['block5']): ?>
    <div class="row">
      <?php if ($content['block4']): ?>
        <div class="col-md-8">
          <?php print $content['block4']; ?>
        </div>
      <?php endif; ?>
      <?php if ($content['block5']): ?>
        <div class="col-md-4">
          <?php print $content['block5']; ?>
        </div>
      <?php endif; ?>
    </div>
  <?php endif; ?>


  <?php if ($content['bottom']): ?>
    <div class="panel-main">
      <?php if ($content['bottom']): ?>
        <div class="row row-bottom">
          <?php print $content['bottom']; ?>
        </div>
      <?php endif; ?>
    </div>
  <?php endif; ?>
  </div>
</div>