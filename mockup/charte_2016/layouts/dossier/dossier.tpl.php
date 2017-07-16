<div class="<?php print $classes ?>" <?php if (!empty($css_id)) { print "id=\"$css_id\""; } ?>>

  <?php if ($content['block1']): ?>
    <div class="row">
      <div class="col-md-12">
        <?php print $content['block1']; ?>
      </div>
    </div>
  <?php endif; ?>
  
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
  
  <?php if ($content['block6']): ?>
    <div class="row">
      <div class="col-md-12">
        <?php print $content['block6']; ?>
      </div>
    </div>
  <?php endif; ?>
  
</div>
